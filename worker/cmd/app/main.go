package main

import (
	"context"
	"log"
	"worker/internal/config"
	"worker/internal/processor"
	"worker/internal/redis"
)

func main() {
	conf, err := config.Load()

	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}

	ctx := context.Background()
	client, err := redis.NewClient(ctx, conf.RedisConfig)

	if err != nil {
		log.Fatalf("Failed to create client: %s\n", err)
	}

	subscriber, err := redis.NewSubscriber(ctx, client, conf.RedisConfig.Channels...)

	if err != nil {
		log.Fatalf("Failed to create subscriber: %s\n", err)
	}

	watcher := processor.NewWatcher(subscriber, client, conf.RedisConfig.Prefix)

	watcher.Run(ctx)
}
