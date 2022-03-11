package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wiisportsresort/chat/server/ent"
	privmsg "github.com/wiisportsresort/chat/server/ent/privatemessage"
	"github.com/wiisportsresort/chat/server/ent/user"
)

func NewGetPrivateMessageRecipients(ctx context.Context, db *ent.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		getPrivateMessageRecipents(ctx, db, c)
	}
}

func getPrivateMessageRecipents(ctx context.Context, db *ent.Client, c *gin.Context) {
	fromId := c.Param("from")

	recipents, err := db.PrivateMessage.Query().
		Where(privmsg.HasAuthorWith(user.ID(fromId))).
		Order(ent.Desc(privmsg.FieldCreatedAt)).
		QueryAuthor().
		All(ctx)

	if err != nil {
		httpError(c, http.StatusInternalServerError, err)
		return
	}

	if len(recipents) == 0 {
		c.JSON(http.StatusOK, ServerHttpMessage{
			Type: "private_messages",
			Data: ServerPrivateMessageRecipientList{
				Recipents: []User{},
			},
		})
		return
	}

	var usersToSend []User
	userIds := make(map[string]bool)

	for _, recipent := range recipents {
		if userIds[recipent.ID] {
			continue
		}
		userIds[recipent.ID] = true
		usersToSend = append(usersToSend, User{
			Id:        recipent.ID,
			Username:  recipent.Username,
			AvatarUrl: recipent.AvatarURL,
		})
	}

	c.JSON(http.StatusOK, ServerHttpMessage{
		Type: "private_message_reciptients",
		Data: ServerPrivateMessageRecipientList{
			Recipents: usersToSend,
		},
	})
}
