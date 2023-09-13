package gateway

import (
	"context"
	"gRPC-chat/api/domain"
	"gRPC-chat/api/driver"
	"gRPC-chat/util"
	"log"
)

type ChatGateway struct {
	driver driver.WebSocketDriver
	redisDriver driver.PubSub
}

func NewChatGateway(driver driver.WebSocketDriver, redisdriver driver.PubSub) *ChatGateway {
	return &ChatGateway{driver: driver, redisDriver: redisdriver}
}

func (g *ChatGateway) Publish(ctx context.Context, msg *domain.Message) *domain.MessageStatus {
	driverMessage := driver.Message{
		Body: msg.Body.Value,
		SendUser: msg.SendUser.Value,
		SendAt: msg.SendAt.ToString(),
		Channel: msg.Channel.Value,
	}
	status := g.redisDriver.Publish(&driverMessage)
	return &domain.MessageStatus{
		Value: status.Code.String(),
		Code: int(status.Code.Val()),
	}
}

func (g *ChatGateway) Subscribe(ctx context.Context, msg *domain.Message) *domain.MessageStatus {
	return &domain.MessageStatus{}
}

func (g *ChatGateway) Broardcast(client *domain.Client, msg *domain.Message) {
	m := driver.Message{
		Body: msg.Body.Value,
		SendUser: msg.SendUser.Value,
		SendAt: msg.SendAt.ToString(),
	}
	g.driver.Broardcast(client.Ws,&m)
}

func (g ChatGateway) ReceiveMessage(client domain.Client) (*domain.Message, error) {
	msg, err := g.driver.ReceiveMessage(client.Ws)
	log.Println("[DEBUG]",util.Totimestamp(msg.SendAt))

	if err != nil {
		return nil, err
	}
	return &domain.Message{
		Body: domain.MessageBody{Value: msg.Body},
		SendUser: domain.SendUser{Value: msg.SendUser},
		SendAt: domain.SendAt{Value: util.Totimestamp(msg.SendAt)},
	}, nil
} 