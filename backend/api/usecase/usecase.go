package usecase

import (
	"context"
	"gRPC-chat/api/domain"
)

//go:generate moq -out InputPort_moq.go . InputPort
type InputPort interface {
	Publish(context.Context, domain.Message) domain.MessageStatus
	Subscribe(context.Context, domain.Message) domain.MessageStatus
}

//go:generate moq -out OutputPort_moq.go . OutputPort
type OutputPort interface {
	SendMessage(context.Context, domain.MessageBody) domain.MessageStatus
	ReceiveMessage(context.Context) domain.Message
}

type UsecaseProvider struct {
	port InputPort
}

func (u *UsecaseProvider) SendMessage(ctx context.Context, msg domain.Message) domain.MessageStatus {
	status :=u.port.Publish(ctx,msg)
	return status
}

func (u *UsecaseProvider) Receivemessage(ctx context.Context) domain.Message {
	return domain.Message{}
}