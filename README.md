# gRPC-chat-go

## Run

~~~bash
go run cmd/main.go
~~~

## Access gRPC

~~~bash
grpcurl -plaintext localhost:8080 list
grpcurl -plaintext localhost:8080 list [サービス名]
grpcurl -plaintext -d '[メッセージ]' localhost:8080 [メソッド名]
~~~
