package driver

import (
	"log"

	"golang.org/x/net/websocket"
)

type WebSocketDriver interface {
	Broardcast(*websocket.Conn,*Message)
	ReceiveMessage(*websocket.Conn) (*Message,error)
}

type Message struct {
	Body     string `json:"msgBody,omitempty"`
	SendUser string `json:"sender,omitempty"`
	SendAt   string `json:"sendAt,omitempty"`
	Channel  string `json:"channel,omitempty"`
}

type WebSocketDriverImpl struct {}

func NewWebSocketDriver() WebSocketDriver {
	return &WebSocketDriverImpl{}
}

func (d *WebSocketDriverImpl) Broardcast(ws *websocket.Conn, msg *Message) {
	websocket.JSON.Send(ws,msg)
}

func (d *WebSocketDriverImpl) ReceiveMessage(ws *websocket.Conn) (*Message, error) {
	msg := Message{}
	err := websocket.JSON.Receive(ws,&msg)
	log.Printf("[INFO] driver receive message : %v", msg)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}