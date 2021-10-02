package model

import "github.com/gorilla/websocket"

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
	UserId int  `json:"user_id"`
	User   User `gorm:"foreignKey:UserId"`
}

type Client struct {
	ID   int
	Conn *websocket.Conn
	Pool *Pool
}

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}
