package main

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func broadcast(msg SocketMessage, sender *Client) {
	for _, client := range clients {
		client.Connection.WriteJSON(msg)
	}
}

func readData(msg SocketMessage, out interface{}) (err error) {
	data, err := json.Marshal(msg.Data)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &out)
	if err != nil {
		return
	}
	return
}

func httpError(c *gin.Context, code int, err error) {
	c.AbortWithStatusJSON(code, SocketMessage{
		Type: "error",
		Data: ServerError{
			Error: err.Error(),
		},
	})
}

func wsError(client *Client, err error) {
	log.Println(err.Error())
	client.Connection.WriteJSON(SocketMessage{
		Type: "error",
		Data: ServerError{
			Error: err.Error(),
		},
	})
}

const (
	OrderModeBefore = "before"
	OrderModeAfter  = "after"
	OrderModeNewest = "newest"
)

type Pagination struct {
	BeforeId  *string
	AfterId   *string
	Limit     int
	OrderMode string
}

func getPagination(c *gin.Context) (pagination Pagination, err error) {
	pagination = Pagination{}

	before_id := c.Query("before")
	after_id := c.Query("after")
	limit_str := c.Query("limit")

	if before_id != "" && after_id != "" {
		err = errors.New("cannot specify both before and after")
		return
	}

	pagination.BeforeId = &before_id
	pagination.AfterId = &after_id

	if before_id != "" {
		pagination.OrderMode = OrderModeBefore
	} else if after_id != "" {
		pagination.OrderMode = OrderModeAfter
	} else {
		pagination.OrderMode = OrderModeNewest
	}

	if limit_str == "" {
		limit_str = "200"
	}

	pagination.Limit, err = strconv.Atoi(limit_str)
	if err != nil {
		err = errors.New("invalid limit")
		return
	}

	if pagination.Limit < 1 || pagination.Limit > 200 {
		err = errors.New("limit must be between 1 and 200")
		return
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
