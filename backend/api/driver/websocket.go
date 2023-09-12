package driver

import (
	"golang.org/x/net/websocket"
)

type WebSocketDriver interface {
	Broardcast(*websocket.Conn,*Message)
	ReceiveMessage(*websocket.Conn) *Message
}

type Message struct {
	Body string
	SendUser string
	SendAt string
}

type WebSocketDriverImpl struct {}

func NewWebSocketDriver() *WebSocketDriverImpl {
	return &WebSocketDriverImpl{}
}

func (d *WebSocketDriverImpl) Broardcast(ws *websocket.Conn, msg *Message) {
	websocket.JSON.Send(ws,msg)
}

func (d *WebSocketDriverImpl) ReceiveMessage(ws *websocket.Conn) *Message {
	msg := Message{}
	websocket.JSON.Receive(ws,&msg)

	return &msg
}