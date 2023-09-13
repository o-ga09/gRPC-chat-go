package driver

import (
	"time"

	"github.com/go-redis/redis"
)

type PubSub interface {
	Publish(*Message) *Status
	Subscribe(*Message) *Subscriber
}

func NewRedis() *redis.Client {
	redisdb := redis.NewClient(&redis.Options{
		Addr:         ":6379",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	return redisdb
}

type PubSubImpl struct {
	redisConn *redis.Client
}

type Status struct {
	Code *redis.IntCmd
}

type Subscriber struct {
	Sub *redis.PubSub
}

func  NewPubSubImpl(redisConn *redis.Client) PubSub {
	return &PubSubImpl{redisConn}
}

func (r *PubSubImpl) Publish(msg *Message) *Status {
	status := r.redisConn.Publish(msg.Channel,msg)

	return &Status{status}
}

func (r *PubSubImpl) Subscribe(msg *Message) *Subscriber {
	subscriber := r.redisConn.Subscribe(msg.Channel)

	return &Subscriber{subscriber}
}	