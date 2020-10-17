package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

var (
	rdb *redis.Client
)

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ENDPOINT"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func main() {
	ctx := context.Background()
	pubsub := rdb.Subscribe(ctx, os.Getenv("REDIS_METRICS_CHANNEL"))
	ch := pubsub.Channel()
	for msg := range ch {
		log.Println(msg.Channel, msg.Payload)
	}
}
