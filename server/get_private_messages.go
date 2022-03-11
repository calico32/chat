package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wiisportsresort/chat/server/ent"
	privmsg "github.com/wiisportsresort/chat/server/ent/privatemessage"
	"github.com/wiisportsresort/chat/server/ent/user"
)

func NewGetPrivateMessages(ctx context.Context, db *ent.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		getPrivateMessages(ctx, db, c)
	}
}

func getPrivateMessages(ctx context.Context, db *ent.Client, c *gin.Context) {
	from_id := c.Param("from")
	to_id := c.Param("to")
	if from_id == to_id {
		httpError(c, http.StatusBadRequest, errors.New(":from and :to cannot be the same"))
		return
	}

	pagination, err := getPagination(c)
	if err != nil {
		httpError(c, http.StatusBadRequest, err)
		return
	}

	beforeId := pagination.BeforeId
	afterId := pagination.AfterId
	limit := pagination.Limit
	orderMode := pagination.OrderMode

	messages, err := db.PrivateMessage.Query().
		Where(privmsg.HasAuthorWith(user.ID(from_id))).
		Where(privmsg.HasRecipientWith(user.ID(to_id))).
		Order(ent.Desc(privmsg.FieldCreatedAt)).
		All(ctx)

	if err != nil {
		httpError(c, http.StatusInternalServerError, err)
		return
	}

	if len(messages) == 0 {
		c.JSON(http.StatusOK, ServerHttpMessage{
			Type: "private_messages",
			Data: ServerPrivateMessageList{
				BeforeId: beforeId,
				AfterId:  afterId,
				Count:    0,
				Messages: []ServerPrivateMessage{},
			},
		})
		return
	}

	var offset int
	if orderMode != OrderModeNewest {
		offset = -1
		for i, message := range messages {
			if message.ID == *beforeId {
				offset = i
				break
			}
		}
		if offset == -1 {
			var err error

			if orderMode == OrderModeBefore {
				err = errors.New(":before not found")
			} else {
				err = errors.New(":after not found")
			}

			httpError(c, http.StatusBadRequest, err)
			return
		}
	}

	var messagesToSend []*ent.PrivateMessage

	switch orderMode {
	case OrderModeNewest:
		messagesToSend = messages[:limit]
	case OrderModeBefore:
		messagesToSend = messages[offset+1 : offset+1+limit]
	case OrderModeAfter:
		messagesToSend = messages[offset-limit : offset]
	}

	var serverMessages []ServerPrivateMessage

	for _, message := range messagesToSend {
		author := message.QueryAuthor().OnlyX(ctx)
		recipient := message.QueryRecipient().OnlyX(ctx)
		parentId := message.QueryParent().OnlyIDX(ctx)

		serverMessages = append(serverMessages, ServerPrivateMessage{
			Id:      message.ID,
			Content: message.Content,
			Author: User{
				Id:        author.ID,
				Username:  author.Username,
				AvatarUrl: author.AvatarURL,
			},
			Recipient: User{
				Id:        recipient.ID,
				Username:  recipient.Username,
				AvatarUrl: recipient.AvatarURL,
			},
			ParentId: &parentId,
		})
	}

	c.JSON(http.StatusOK, ServerHttpMessage{
		Type: "private_messages",
		Data: ServerPrivateMessageList{
			BeforeId: beforeId,
			AfterId:  afterId,
			Count:    0,
			Messages: serverMessages,
		},
	})
}
