package domain

import (
	"log"
	"time"

	"golang.org/x/net/websocket"
)

type Room struct {
	Clients map[*Client]bool
	RegisterCh chan *Client
	UnRegisterCh chan *Client
	BroardcastCh chan []byte
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
		select {
		case client := <- r.RegisterCh:
			r.Register(client)
		case client := <- r.UnRegisterCh:
			r.UnRegister(client)
		case msg := <- r.BroardcastCh:
			r.Broardcast(msg)
		}
	}
}

func (r *Room) Register(client *Client) {
	r.Clients[client] = true
}

func (r *Room) UnRegister(client *Client) {
	delete(r.Clients,client)
}

func (r * Room) Broardcast(msg []byte) {
	for c := range r.Clients {
		c.sendCh <- msg
	}
}

func (c *Client) Read(broardcastcha <- chan []byte, unregister <- chan *Client) {
	defer c.ws.Close()

	for {
		var msg []byte
		_, err := c.ws.Read(msg)
		if err != nil {
			// slogに変更
			log.Printf("[ERR] can not read message: %v",err)
			if !c.ws.IsClientConn() {
				log.Printf("[INFO] unexpected server connection close")
			}
			break
		}

		c.sendCh <- msg
	}
}

func (c *Client) Write() {
	defer c.ws.Close()

	for {
		msg := <- c.sendCh

		_, err := c.ws.Write(msg)
		if err != nil {
			// slogに変更
			log.Printf("[ERR] can not write message: %v",err)
			return
		}

		if err := c.ws.Close(); err != nil {
			return
		}
	}
}

func (c *Client) disconnect(unregister chan<- *Client) {
	unregister <- c
	c.ws.Close()
}