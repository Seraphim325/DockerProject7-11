package processor

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"
	"worker/internal/helper"
	redis_worker "worker/internal/redis"

	"github.com/redis/go-redis/v9"
)

type Watcher struct {
	subscriber *redis_worker.Subscriber
	client     *redis.Client
	prefix     string
}

func NewWatcher(subscriber *redis_worker.Subscriber, client *redis.Client, prefix string) *Watcher {
	return &Watcher{
		subscriber: subscriber,
		client:     client,
		prefix:     prefix,
	}
}

func (w *Watcher) Run(ctx context.Context) error {
	defer w.subscriber.Close()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg, ok := <-w.subscriber.Messages():
			if !ok {
				return nil
			}
			w.handleMessage(ctx, msg)
		}
	}

}

func (w *Watcher) handleMessage(ctx context.Context, msg *redis.Message) {
	content := msg.Payload

	if !strings.HasPrefix(content, w.prefix) {
		return
	}

	if _, err := w.client.Del(ctx, content).Result(); err != nil {
		log.Printf("Failed to delete key %s: %s\n", content, err)
	}

	key := content[6:]
	keyConverted, err := strconv.Atoi(key)

	if err != nil {
		log.Printf("Failed to convert key %s to value: %s\n", key, err)
		return
	}

	if err := w.client.Set(ctx, key, helper.Fib(keyConverted), 24*time.Hour).Err(); err != nil {
		log.Printf("Failed to set value for key %s: %s\n", key, err)
	}
}
