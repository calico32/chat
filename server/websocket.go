package main

import (
	"context"
	"crypto"
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	gonanoid "github.com/matoous/go-nanoid"
	"github.com/wiisportsresort/chat/server/ent"
	entMessage "github.com/wiisportsresort/chat/server/ent/message"
	entUser "github.com/wiisportsresort/chat/server/ent/user"
)

func NewWebsocketHandler(ctx context.Context, db *ent.Client, upgrader *websocket.Upgrader) func(c *gin.Context) {
	return func(c *gin.Context) {
		handleWebsocket(ctx, db, upgrader, c)
	}
}

func handleWebsocket(ctx context.Context, db *ent.Client, upgrader *websocket.Upgrader, c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, SocketMessage{
			Type: "error",
			Data: ServerError{
				Error: err.Error(),
			},
		})
		return
	}

	id := gonanoid.MustID(24)
	client := &Client{
		Connection: ws,
		Id:         id,
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
		ws.Close()
		delete(clients, client.Id)
	}()

	for {
		var msg SocketMessage
		err := ws.ReadJSON(&msg)
		if err != nil {
			wsError(client, err)
			break
		}

		if msg.Type == "login" {
			var login ClientLogin
			err = readData(msg, &login)
			if err != nil {
				wsError(client, err)
				continue
			}

			newId, err := generateId(login.User)
			if err != nil {
				wsError(client, err)
				continue
			}

			user, err := findUser(login.User, db, ctx)
			if err != nil {
				wsError(client, err)
				continue
			}

			id = newId
			client.Id = newId
			client.EntUser = user
			clients[newId] = client

			client.User = &User{
				Id:         newId,
				Username:   user.Username,
				AvatarUrl:  user.AvatarURL,
				SecretHash: user.SecretHash,
			}

			message := SocketMessage{
				Type: "login",
				Data: ServerLogin{
					User: client.User,
				},
			}

			client.Connection.WriteJSON(message)
		}

		if client.User == nil {
			wsError(client, errors.New("User not logged in"))
			break
		}

		if msg.Type == "message" {
			var message ClientMessage
			err = readData(msg, &message)
			if err != nil {
				wsError(client, err)
				continue
			}

			isAdmin := false
			for _, admin := range admins {
				if client.User.Id == admin {
					isAdmin = true
				}
			}

			if isAdmin {
				if strings.HasPrefix(message.Content, "!!") {
					command := strings.TrimPrefix(message.Content, "!!")
					switch command {
					case "clear":
						_, err := db.Message.Delete().Exec(ctx)
						if err != nil {
							wsError(client, err)
							continue
						}
						broadcast(SocketMessage{
							Type: "clear",
							Data: nil,
						}, client)
						continue
					case "reload":
						broadcast(SocketMessage{
							Type: "reload",
							Data: nil,
						}, client)
					default:
						wsError(client, errors.New("unknown command"))
					}

					continue
				}
			}

			messageId := gonanoid.MustID(24)

			msgBuilder := db.Message.Create().
				SetID(messageId).
				SetAuthor(client.EntUser).
				SetContent(message.Content).
				SetCreatedAt(time.Now())

			if message.ParentId != nil {
				parent, err := db.Message.Query().Where(entMessage.IDEQ(*message.ParentId)).Only(ctx)

				if err != nil {
					wsError(client, errors.New("parent message not found"))
					continue
				}

				msgBuilder = msgBuilder.SetParent(parent)
			}

			msg, err := msgBuilder.Save(ctx)
			if err != nil {
				wsError(client, err)
				continue
			}

			serverMessage := SocketMessage{
				Type: "message",
				Data: ServerMessage{
					Id:        msg.ID,
					Content:   msg.Content,
					Author:    *client.User,
					ParentId:  message.ParentId,
					CreatedAt: uint64(msg.CreatedAt.UnixMilli()),
				},
			}

			broadcast(serverMessage, client)
		}

		if msg.Type == "privateMessage" {
			var privateMessage ClientPrivateMessage
			err = readData(msg, &privateMessage)
			if err != nil {
				wsError(client, err)
				continue
			}

			messageId := gonanoid.MustID(24)

			serverMessage := SocketMessage{
				Type: "privateMessage",
				Data: ServerPrivateMessage{
					Id:       messageId,
					Content:  privateMessage.Content,
					Author:   *client.User,
					ParentId: privateMessage.ParentId,
				},
			}

			if client, ok := clients[privateMessage.ToId]; ok {
				client.Connection.WriteJSON(serverMessage)
			} else {
				wsError(client, errors.New("User not found"))
			}
		}
	}
}

func generateId(user UserNoId) (string, error) {
	hashComponent := user.Username + ":"
	if user.SecretHash != nil {
		if len(*user.SecretHash) < 8 {
			return "", errors.New("secret hash too short")
		}
		hashComponent += *user.SecretHash
	}

	hash := crypto.SHA1.New()
	_, err := hash.Write([]byte(hashComponent))
	if err != nil {
		return "", err
	}
	b64Hash := base64.URLEncoding.EncodeToString(hash.Sum(nil))
	return b64Hash, nil
}

func findUser(user UserNoId, db *ent.Client, ctx context.Context) (*ent.User, error) {
	id, err := generateId(user)
	if err != nil {
		return nil, err
	}

	existing, err := db.User.Query().Where(entUser.IDEQ(id)).All(ctx)

	if err != nil {
		return nil, err
	}

	if len(existing) > 0 {
		dbUser := existing[0]
		if dbUser.AvatarURL != user.AvatarUrl {
			if user.AvatarUrl != nil {
				dbUser, err = dbUser.Update().SetAvatarURL(*user.AvatarUrl).Save(ctx)
			} else {
				dbUser, err = dbUser.Update().ClearAvatarURL().Save(ctx)
			}

			if err != nil {
				return nil, err
			}
		}

		return dbUser, nil
	}

	userBuilder := db.User.Create().
		SetID(id).
		SetUsername(user.Username).
		SetNillableAvatarURL(user.AvatarUrl).
		SetNillableSecretHash(user.SecretHash)

	return userBuilder.Save(ctx)
}
