package driver

import (
	"time"

	"github.com/go-redis/redis"
)

type PubSub interface {
	Publish(Message) Status
	Subcribe(Message) Subscriber
}

type Message struct {
	Body string
	SendUser string
	SendAt time.Time
}

type Status struct {
	Code *redis.IntCmd
}

type Subscriber struct {
	Sub *redis.PubSub
}