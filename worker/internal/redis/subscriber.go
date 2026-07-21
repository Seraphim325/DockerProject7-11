package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Subscriber struct {
	pubsub *redis.PubSub
}

func NewSubscriber(ctx context.Context, client *redis.Client, channels ...string) (*Subscriber, error) {
	pubsub := client.PSubscribe(ctx, channels...)

	if _, err := pubsub.Receive(ctx); err != nil {
		return nil, err
	}

	return &Subscriber{pubsub: pubsub}, nil
}

func (s *Subscriber) Messages() <-chan *redis.Message {
	return s.pubsub.Channel()
}

func (s *Subscriber) Close() error {
	return s.pubsub.Close()
}
