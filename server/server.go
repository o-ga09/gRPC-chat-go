package server

import (
	"fmt"
	chatpb "gRPC-chat/pkg/grpc"
	"gRPC-chat/server/handler"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Run() {
	port := 8080
	l, err := net.Listen("tcp",fmt.Sprintf(":%d",port))
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	chatpb.RegisterMessagingServiseServer(server,handler.NewChatServer())
	reflection.Register(server)

	go func() {
		log.Printf("server started on :%d",port)
		server.Serve(l)
	} ()

	quit := make(chan os.Signal,1)
	signal.Notify(quit,os.Interrupt)
	<- quit
	log.Printf("stopping server ...")
	server.GracefulStop()
}