package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wiisportsresort/chat/server/ent"
	"github.com/wiisportsresort/chat/server/ent/message"
)

func NewGetMessages(ctx context.Context, db *ent.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		getMessages(ctx, db, c)
	}
}

func getMessages(ctx context.Context, db *ent.Client, c *gin.Context) {
	channel := c.Param("channel")

	if channel != "public" {
		httpError(c, http.StatusBadRequest, errors.New("invalid channel"))
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

	messages, err := db.Message.Query().
		Order(ent.Desc(message.FieldCreatedAt)).
		All(ctx)

	if err != nil {
		httpError(c, http.StatusInternalServerError, err)
		return
	}

	if len(messages) == 0 {
		c.JSON(http.StatusOK, ServerHttpMessage{
			Type: "messages",
			Data: ServerMessageList{
				BeforeId: beforeId,
				AfterId:  afterId,
				Count:    0,
				Messages: []ServerMessage{},
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

	var messagesToSend []*ent.Message

	switch orderMode {
	case OrderModeNewest:
		messagesToSend = messages[:min(len(messages), limit)]
	case OrderModeBefore:
		messagesToSend = messages[max(0, offset+1):min(len(messages), offset+1+limit)]
	case OrderModeAfter:
		messagesToSend = messages[max(0, offset-limit):min(len(messages), offset)]
	}

	var serverMessages []ServerMessage

	for _, message := range messagesToSend {
		parent := message.QueryParent().AllX(ctx)
		author := message.QueryAuthor().OnlyX(ctx)

		var parentId *string
		if len(parent) > 0 {
			parentId = &parent[0].ID
		}

		serverMessages = append(serverMessages, ServerMessage{
			Id:       message.ID,
			ParentId: parentId,
			Author: User{
				Id:        author.ID,
				Username:  author.Username,
				AvatarUrl: author.AvatarURL,
			},
			Content:   message.Content,
			CreatedAt: uint64(message.CreatedAt.UnixMilli()),
		})
	}

	c.JSON(http.StatusOK, ServerHttpMessage{
		Type: "messages",
		Data: ServerMessageList{
			BeforeId: beforeId,
			AfterId:  afterId,
			Count:    0,
			Messages: serverMessages,
		},
	})

}
