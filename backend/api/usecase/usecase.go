package usecase

import (
	"context"
	"errors"
	"gRPC-chat/api/domain"
	"io"
	"log"

	"golang.org/x/net/websocket"
)

//go:generate moq -out InputPort_moq.go . InputPort
type InputPort interface {
	Publish(context.Context, *domain.Message) *domain.MessageStatus
	Subscribe(context.Context, *domain.Message) *domain.MessageStatus
	Broardcast(client *domain.Client,msg *domain.Message)
	ReceiveMessage(domain.Client) (*domain.Message, error)
}

type UsecaseProvider struct {
	port InputPort
	Room *domain.Room
}

func NewUsecase(port InputPort) *UsecaseProvider {
	return &UsecaseProvider{
		port: port,
		Room: domain.NewRoom(),
	}
}

func (u *UsecaseProvider) ReceiveAndSend(client *domain.Client) error {
	defer func() {
		recover()
	}()
	for {
		msg, err := u.port.ReceiveMessage(*client)
		if err != nil {
			if errors.Is(err,io.EOF) {
				continue
			}
			return err
		}

		log.Printf("[INFO] receive message :  %v",msg)

		u.Room.Message <- msg
	}
}

func (u *UsecaseProvider) ReceiveMessage() {
	for {
		message := <-u.Room.Message
		u.SendMessage(message)
	}
}

func (u *UsecaseProvider) Register(client *domain.Client) {
	u.Room.Clients[client] = true
}

func (u *UsecaseProvider) UnRegister(client *domain.Client) {
	delete(u.Room.Clients,client)
}

func (u *UsecaseProvider) SendMessage(msg *domain.Message) {
	for c := range u.Room.Clients {
		u.port.Broardcast(c,msg)
	}
}

func (u *UsecaseProvider) NewClient(ctx context.Context) *domain.Client {
	v := ctx.Value("wsConn")
	ws, ok := v.(*websocket.Conn)
	if !ok {
		log.Printf("[ERR] invalid websocket connection type")
		panic("err")
	}
	return domain.NewClient(ws)
}