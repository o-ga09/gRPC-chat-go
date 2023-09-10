package grpc

import (
	"context"
	chatpb "gRPC-chat/pkg/grpc"

	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatService struct {
	chatpb.UnimplementedMessagingServiseServer
}

func NewChatServer() *ChatService {
	return &ChatService{}
}

func(s *ChatService) HealthCheck(ctx context.Context, empty *emptypb.Empty) (*chatpb.Status, error) {
	return &chatpb.Status{Msg: "OK"}, nil
}

func(s *ChatService) SendMessage(ctx context.Context, req *chatpb.MsgRequest) (*chatpb.MsgResponse, error) {
	res := chatpb.MsgResponse{
		MsgBody: "Hello World !",
		SendTimestamp: req.SendTimestamp,
	}
	return &res, nil
}

func(s *ChatService) ReceiveMessage(req *chatpb.MsgRequest, stream chatpb.MessagingServise_ReceiveMessageServer) error {
	chatNum := 2
	for i := 0; i < chatNum; i++ {
		stream.Send(&chatpb.MsgResponse{
			MsgBody: req.GetMsgBody(),
			SendTimestamp: req.GetSendTimestamp(),
		})
	}
	return nil
} 