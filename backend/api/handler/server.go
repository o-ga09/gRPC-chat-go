package handler

import (
	"fmt"
	"gRPC-chat/api/domain"
	gp "gRPC-chat/api/handler/grpc"
	ws "gRPC-chat/api/handler/websocket"
	chatpb "gRPC-chat/pkg/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/net/websocket"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run() {
	grpc_port := 8080
	grpc_listener, err := net.Listen("tcp",fmt.Sprintf(":%d",grpc_port))
	if err != nil {
		log.Fatal(err)
	}

	websocket_port := 8081
	websocket_listener, err := net.Listen("tcp",fmt.Sprintf(":%d",websocket_port))
	if err != nil {
		log.Fatal(err)
	}

	grpc_server := grpc.NewServer()
	chatpb.RegisterMessagingServiseServer(grpc_server,gp.NewChatServer())
	reflection.Register(grpc_server)

	room := domain.NewRoom()
	websocket_handler := ws.NewWebSocketHandler(room)
	websocket_server := http.Server{
		Handler: websocket.Handler(websocket_handler.Handler),
	}
	// WebSocketハンドラを設定
	// http.Handle("/ws",websocket_server.Handler)
	http.Handle("/ws", websocket.Handler(func(ws *websocket.Conn) {
		websocket_handler.Handler(ws)
	}))

	go func() {
		log.Printf("websocket server started on :%d",websocket_port)
		websocket_server.Serve(websocket_listener)
	} ()

	go func() {
		log.Printf("gRPC server started on :%d",grpc_port)
		grpc_server.Serve(grpc_listener)
	} ()

	go func() {
		log.Printf("Chat Room goroutine start ...")
		room.Run()
	} ()

	quit := make(chan os.Signal,1)
	signal.Notify(quit,os.Interrupt)
	<- quit
	log.Printf("stopping server ...")
	grpc_server.GracefulStop()
}