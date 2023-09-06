package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	// server.Run()
	redisdb := redis.NewClient(&redis.Options{
		Addr:         ":6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	pubsub := redisdb.Subscribe("testchannel")

	// Wait for confirmation that subscription is created before publishing anything.
	_, err := pubsub.Receive()
	if err != nil {
		panic(err)
	}

	// Go channel which receives messages.
	subChan := make(chan string)
	go func(){
		for {
			ch := pubsub.Channel()
			// Consume messages.
			for msg := range ch {
				subChan <- msg.Payload
			}
		}
	}()

	// Go channel which publish messages.
	pubChan := make(chan string)
	go func() {
		b := bufio.NewReader(os.Stdin)
		for {
			fmt.Printf(">")
			line, _, err := b.ReadLine()
			if err != nil {
				log.Println(err)
				pubChan <- "/exit"
				return
			}
			pubChan <- string(line)
		}
	} ()

	chatting := true
	for chatting {
		select {
		case msg := <- subChan:
			fmt.Println(msg)
			fmt.Printf(">")
		case input := <- pubChan:
			if input == "/exit" {
				chatting = false
			} else {
				redisdb.Publish("testchannel",input)
			}
		default:
			time.Sleep(1* time.Millisecond)
		}
	}

	redisdb.Publish("testchannel","user has left")

	time.AfterFunc(time.Second, func() {
		// When pubsub is closed channel is closed too.
		_ = pubsub.Close()
	})
}