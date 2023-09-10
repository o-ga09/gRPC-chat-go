package wensocket

import (
	"gRPC-chat/api/domain"
	"log"
	"net/http"

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
	defer ws.Close()

	log.Printf("[INFO] user enter")
	client := domain.NewClient(ws)
	go client.Read(h.room.BroardcastCh,h.room.RegisterCh)
	go client.Write()
	h.room.RegisterCh <- client
}

func (h *WebsocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}