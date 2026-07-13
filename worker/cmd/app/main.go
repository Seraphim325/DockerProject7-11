package app

import (
	"context"
	"worker/internal/config"
	"worker/internal/redis"
)

func main() {
	config.LoadConfig()
	redisConf := config.LoadRedisConfig()
	ctx := context.Background()
	client := redis.EstablishConnection(ctx, redisConf)

	redis.Subscribe(ctx, client)
}
