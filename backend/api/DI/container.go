package di

import (
	"gRPC-chat/api/driver"
	"gRPC-chat/api/gateway"
	"gRPC-chat/api/handler/websocket"
	"gRPC-chat/api/usecase"
)

func DiContainer() *websocket.WebsocketHandler {
	driver := driver.NewWebSocketDriver()
	gateway := gateway.NewChatGateway(driver)
	usecase := usecase.NewUsecase(gateway)
	handler := websocket.NewWebSocketHandler(usecase)

	return handler
}