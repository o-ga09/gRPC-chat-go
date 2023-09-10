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
	Message chan string
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
	ws *websocket.Conn
	sendCh chan []byte
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
		Message: make(chan string),
	}
}

func NewClient(ws *websocket.Conn) *Client {
	return &Client{
		ws: ws,
		sendCh: make(chan []byte),
	}
}

func (r *Room) Run() {
	for {
		message := <-r.Message
		r.Broardcast(message)
	}
}

func (r *Room) Register(client *Client) {
	r.Clients[client] = true
}

func (r *Room) UnRegister(client *Client) {
	delete(r.Clients,client)
}

func (r * Room) Broardcast(msg string) {
	for c := range r.Clients {
		websocket.Message.Send(c.ws,msg)
	}
}