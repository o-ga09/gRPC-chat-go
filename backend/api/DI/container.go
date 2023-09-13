package di

import (
	"gRPC-chat/api/driver"
	"gRPC-chat/api/gateway"
	"gRPC-chat/api/handler/websocket"
	"gRPC-chat/api/usecase"
)

func DiContainer() *websocket.WebsocketHandler {
	redisConn := driver.NewRedis()
	redisDriver := driver.NewPubSubImpl(redisConn)
	driver := driver.NewWebSocketDriver()
	gateway := gateway.NewChatGateway(driver,redisDriver)
	usecase := usecase.NewUsecase(gateway)
	handler := websocket.NewWebSocketHandler(usecase)

	return handler
}