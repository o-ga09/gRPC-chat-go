package driver

import (
	"github.com/go-redis/redis"
)

type PubSub interface {
	Publish(Message) Status
	Subcribe(Message) Subscriber
}

type Status struct {
	Code *redis.IntCmd
}

type Subscriber struct {
	Sub *redis.PubSub
}