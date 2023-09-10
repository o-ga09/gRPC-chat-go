package main

import (
	"gRPC-chat/api/handler"
	"time"
)

type message struct {
	Body string
	SendUser string
	SendAt time.Time
}

func main() {
	handler.Run()
	// defer func(){recover()}()
	// redisdb := redis.NewClient(&redis.Options{
	// 	Addr:         ":6379",
	// 	DialTimeout:  10 * time.Second,
	// 	ReadTimeout:  30 * time.Second,
	// 	WriteTimeout: 30 * time.Second,
	// 	PoolSize:     10,
	// 	PoolTimeout:  30 * time.Second,
	// })

	// pubsub := redisdb.Subscribe("testchannel")

	// buf := bufio.NewReader(os.Stdin)
	// fmt.Printf("Input Your Name > ")
	// username, _, err := buf.ReadLine()
	// if err != nil {
	// 	panic("User name Invalid")
	// }

	// // Go channel which receives messages.
	// subChan := make(chan *message)
	// go func(){
	// 	for {
	// 		msg, err := pubsub.ReceiveMessage()
	// 		m := new(message)
	// 		json.Unmarshal([]byte(msg.Payload),&m)
	// 		if err == nil {
	// 			subChan <- m
	// 		}
	// 	}
	// }()

	// // Go channel which publish messages.
	// pubChan := make(chan string)
	// go func() {
	// 	b := bufio.NewReader(os.Stdin)
	// 	for {
	// 		fmt.Printf(">")
	// 		line, _, err := b.ReadLine()
	// 		if err != nil {
	// 			log.Println(err)
	// 			pubChan <- "/exit"
	// 			return
	// 		}
	// 		pubChan <- string(line)
	// 	}
	// } ()

	// chatting := true
	// for chatting {
	// 	select {
	// 	case msg := <- subChan:
	// 		if msg.SendUser == string(username) {
	// 			continue
	// 		}
	// 		fmt.Printf("%s > %s\n",msg.SendUser,msg.Body)
	// 		fmt.Printf(">")
	// 	case input := <- pubChan:
	// 		if input == "/exit" {
	// 			chatting = false
	// 		} else {
	// 			m := message{
	// 				Body: input,
	// 				SendUser: string(username),
	// 				SendAt: time.Now(),
	// 			}
	// 			j, _ := json.Marshal(m)
	// 			redisdb.Publish("testchannel",j)
	// 		}
	// 	default:
	// 		time.Sleep(1* time.Millisecond)
	// 	}
	// }

	// redisdb.Publish("testchannel","user has left")

	// time.AfterFunc(time.Second, func() {
	// 	// When pubsub is closed channel is closed too.
	// 	_ = pubsub.Close()
	// })
}