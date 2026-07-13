package redis

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"worker/internal/helper"

	"github.com/redis/go-redis/v9"
)

const (
	prefix = "index:"
)

func Subscribe(ctx context.Context, client *redis.Client) {
	pubsub := client.PSubscribe(ctx, "__keyevent@0__:set", "__keyevent@0__:hset")
	defer pubsub.Close()

	if _, err := pubsub.Receive(ctx); err != nil {
		log.Fatalf("Failed to subscribe to redis channel: %s\n", err)
	}

	msgCh := pubsub.Channel()

	for msg := range msgCh {
		key := msg.Payload
		if !strings.HasPrefix(key, prefix) {
			continue
		}

		if _, err := client.Del(ctx, key).Result(); err != nil {
			log.Printf("Failed to delete key %s: %s\n", key, err)
		}

		computedKey := key[6:]
		computedValue, err := strconv.Atoi(computedKey)

		if err != nil {
			log.Printf("Failed to convert value %s: %s\n", computedKey, err)
			continue
		}

		client.Set(ctx, computedKey, helper.Fib(computedValue), 24*time.Hour)

	}
}
