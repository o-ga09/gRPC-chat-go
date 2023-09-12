package websocket

import (
	"context"
	"gRPC-chat/api/domain"
	"gRPC-chat/api/usecase"
	"log"

	"golang.org/x/net/websocket"
)

type WebsocketHandler struct {
	Usecase *usecase.UsecaseProvider
}

func NewWebSocketHandler(usecase *usecase.UsecaseProvider) *WebsocketHandler {
	return &WebsocketHandler{
		Usecase: usecase,
	}
}

func (h *WebsocketHandler) Handler(ctx context.Context, ws *websocket.Conn) {
	defer func () {
		log.Printf("[INFO] websocket connection closed")
		ws.Close()
	}()


	client := h.Usecase.NewClient(ctx)
	h.Usecase.Register(client)

	var msg domain.Message
	for {
		h.Usecase.Receivemessage()
		log.Printf("[INFO] receive message :  %v",msg)

		h.Usecase.Room.Message <- msg
	}
}