package domain

import (
	"time"

	"golang.org/x/net/websocket"
)

type Room struct {
	Clients map[*Client]bool
	RegisterCh chan *Client
	UnRegisterCh chan *Client
	BroardcastCh chan []byte
	Message chan Message
}

type Message struct {
	Body MessageBody
	SendUser SendUser
	SendAt SendAt
}

type MessageBody struct {
	Value string
}

type SendUser struct {
	Value string
}

type SendAt struct {
	Value time.Time
}

type MessageStatus struct {
	Value string
	Code int
}

type Client struct {
	Ws *websocket.Conn
	SendCh chan []byte
}

const (
	SUCCESS = 0
	ERR = 1
)

func NewRoom() *Room {
	return &Room{
		Clients: make(map[*Client]bool),
		RegisterCh: make(chan *Client),
		UnRegisterCh: make(chan *Client),
		BroardcastCh: make(chan []byte),
		Message: make(chan Message),
	}
}

func NewClient(ws *websocket.Conn) *Client {
	return &Client{
		Ws: ws,
		SendCh: make(chan []byte),
	}
}

func (t *SendAt) ToString() string {
	return t.Value.String()
}