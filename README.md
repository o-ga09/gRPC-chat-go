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

## 記事化メモ

### GoでWebsockeを使用してチャットを作ろうとしたら詰まったこと

- 元々参考にしていた記事

https://zenn.dev/tady/articles/adcdc65617ae57

- chatgptに聞いた回答

goでwebsocket通信を使用したチャットを作ってください。
ただし、使用するライブラリは、「"golang.org/x/net/websocket"」でお願いします。
ChatGPT
もちろんです。以下はGo言語を使用してWebSocket通信を利用した簡単なチャットアプリケーションの例です。この例では、golang.org/x/net/websocket パッケージを使用します。

まず、main.go ファイルを作成し、以下のコードを追加します：

~~~go
package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

// チャットルーム構造体
type ChatRoom struct {
	clients map[*websocket.Conn]bool
	messages chan string
}

// メッセージのブロードキャスト
func (c *ChatRoom) broadcast(message string) {
	for client := range c.clients {
		websocket.Message.Send(client, message)
	}
}

// WebSocketハンドラ
func (c *ChatRoom) handleWebSocket(conn *websocket.Conn) {
	c.clients[conn] = true
	defer func() {
		delete(c.clients, conn)
		conn.Close()
	}()

	var message string
	for {
		err := websocket.Message.Receive(conn, &message)
		if err != nil {
			break
		}
		c.messages <- message
	}
}

// メッセージを受信しブロードキャスト
func (c *ChatRoom) processMessages() {
	for {
		message := <-c.messages
		c.broadcast(message)
	}
}

func main() {
	// チャットルームの初期化
	chatRoom := &ChatRoom{
    clients:  make(map[*websocket.Conn]bool),
	messages: make(chan string),
	}
	go chatRoom.processMessages()

	// WebSocketハンドラを設定
	http.Handle("/chat", websocket.Handler(func(ws *websocket.Conn) {
		chatRoom.handleWebSocket(ws)
	}))

	// サーバーを起動
	fmt.Println("WebSocketサーバーが :8080 で起動しました。")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
~~~

このコードでは、WebSocketを使用してクライアント間でメッセージを送受信できるシンプルなチャットルームを作成しています。

以下の手順で動作を確認できます：

ターミナルでプロジェクトのディレクトリに移動し、main.go ファイルを作成します。
上記のコードを main.go ファイルに貼り付けます。
コンソールで go run main.go を実行します。
WebSocketサーバーは localhost:8080 で起動されます。

これで、WebSocketを使用した簡単なチャットアプリケーションが完成しました。この例では、単純なブロードキャストのみを行っていますが、実際のアプリケーションではセキュリティやエラーハンドリングなど、さらなる機能が必要になるかもしれません。
