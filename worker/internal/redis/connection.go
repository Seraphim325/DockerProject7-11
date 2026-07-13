package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func EstablishConnection(ctx context.Context, conf RedisConfig) *redis.Client {
	client, err := NewClient(ctx, conf)

	if err != nil {
		log.Fatalf("Failed to establish connection with redis: %s\n", err)
	}

	return client
}
