package handler

import (
	"context"
	"fmt"
	di "gRPC-chat/api/DI"
	gp "gRPC-chat/api/handler/grpc"
	chatpb "gRPC-chat/pkg/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	websocket_handler := di.DiContainer()
	websocket_server := http.Server{
		Handler: websocket.Handler(func(ws *websocket.Conn) {
			ctx := context.Background()
			ctx = context.WithValue(ctx,"sendAt",time.Now())
			websocket_handler.Handler(ctx,ws)
		}),
	}
	// WebSocketハンドラを設定
	// http.Handle("/ws",websocket_server.Handler)
	http.Handle("/ws",websocket_server.Handler)

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
		websocket_handler.Usecase.ReceiveMessage()
	} ()

	quit := make(chan os.Signal,1)
	signal.Notify(quit,os.Interrupt)
	<- quit
	log.Printf("stopping server ...")
	grpc_server.GracefulStop()
}