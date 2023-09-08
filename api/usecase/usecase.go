package usecase

import "gRPC-chat/api/domain"

type port interface {
	SendMessage(domain.MessageBody) domain.MessageStatus
	ReceiveMessage() domain.MessageBody
}