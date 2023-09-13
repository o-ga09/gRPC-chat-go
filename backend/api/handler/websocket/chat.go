package websocket

import (
	"context"
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

	ctx = context.WithValue(ctx,"wsConn",ws)
	client := h.Usecase.NewClient(ctx)
	h.Usecase.Register(client)
	log.Println("[INFO] user registerd")

	err := h.Usecase.ReceiveAndSend(client)
	log.Printf("[INFO] err : %v", err)
}