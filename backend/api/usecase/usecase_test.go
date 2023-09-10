package usecase

import (
	"context"
	"gRPC-chat/api/domain"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSendMessage(t *testing.T) {
	mock := &InputPortMock{
		PublishFunc: func(contextMoqParam context.Context, message domain.Message) domain.MessageStatus {
			return domain.MessageStatus{}
		},
		SubscribeFunc: func(contextMoqParam context.Context, message domain.Message) domain.MessageStatus {
			return domain.MessageStatus{}
		},
	}

	usecase := UsecaseProvider{mock}
	msg := domain.Message{}
	ctx := context.Background()
	actual := usecase.SendMessage(ctx,msg)
	expected := domain.MessageStatus{}
	assert.Equal(t,expected,actual)
}

func TestReceiveMessage(t *testing.T) {
	mock := &InputPortMock{
		PublishFunc: func(contextMoqParam context.Context, message domain.Message) domain.MessageStatus {
			return domain.MessageStatus{}
		},
		SubscribeFunc: func(contextMoqParam context.Context, message domain.Message) domain.MessageStatus {
			return domain.MessageStatus{}
		},
	}

	usecase := UsecaseProvider{mock}
	ctx := context.Background()
	actual := usecase.Receivemessage(ctx)
	expected := domain.Message{}
	assert.Equal(t,expected,actual)
}