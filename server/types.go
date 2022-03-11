package main

import (
	"github.com/gorilla/websocket"
	"github.com/wiisportsresort/chat/server/ent"
)

type UserNoId struct {
	Username   string  `json:"username"`
	AvatarUrl  *string `json:"avatarUrl"`
	SecretHash *string `json:"secretHash"`
}

type User struct {
	Id         string  `json:"id"`
	Username   string  `json:"username"`
	AvatarUrl  *string `json:"avatarUrl"`
	SecretHash *string `json:"secretHash"`
}

type Client struct {
	Id         string
	Connection *websocket.Conn
	User       *User
	EntUser    *ent.User
}

type SocketMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// client

type ClientLogin struct {
	User UserNoId `json:"user"`
}

type ClientMessage struct {
	Content  string  `json:"content"`
	ParentId *string `json:"parentId"`
}

type ClientPrivateMessage struct {
	Content  string  `json:"content"`
	ToId     string  `json:"toId"`
	ParentId *string `json:"parentId"`
}

// server

type ServerLogin struct {
	User *User `json:"user"`
}

type ServerError struct {
	Error string `json:"error"`
}

type ServerMessage struct {
	Content   string  `json:"content"`
	ParentId  *string `json:"parentId"`
	Id        string  `json:"id"`
	Author    User    `json:"author"`
	CreatedAt uint64  `json:"createdAt"`
}

type ServerPrivateMessage struct {
	Id        string  `json:"id"`
	Content   string  `json:"content"`
	Author    User    `json:"from"`
	Recipient User    `json:"to"`
	ParentId  *string `json:"parentId"`
}

// http

type ServerHttpMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type ServerMessageList struct {
	BeforeId *string         `json:"beforeId"`
	AfterId  *string         `json:"afterId"`
	Count    int             `json:"count"`
	Messages []ServerMessage `json:"messages"`
}

type ServerPrivateMessageList struct {
	BeforeId *string                `json:"beforeId"`
	AfterId  *string                `json:"afterId"`
	Count    int                    `json:"count"`
	Messages []ServerPrivateMessage `json:"messages"`
}

type ServerPrivateMessageRecipientList struct {
	Recipents []User `json:"recipents"`
}
