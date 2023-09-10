package wensocket

import (
	"gRPC-chat/api/domain"
	"log"

	"golang.org/x/net/websocket"
)

type WebsocketHandler struct {
	room *domain.Room
}

func NewWebSocketHandler(room *domain.Room) *WebsocketHandler {
	return &WebsocketHandler{
		room: room,
	}
}

func (h *WebsocketHandler) Handler(ws *websocket.Conn) {
	defer func () {
		log.Printf("[INFO] websocket connection closed")
		ws.Close()
	}()

	client := domain.NewClient(ws)
	h.room.Register(client)

	var msg string
	for {
		err := websocket.Message.Receive(ws, &msg)
		log.Printf("[INFO] receive message :  %v",msg)
		if err != nil {
			break
		}
		h.room.Message <- msg
	}
}