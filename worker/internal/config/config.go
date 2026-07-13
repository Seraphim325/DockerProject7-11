package config

import (
	"fmt"
	"log"
	"os"
	"time"
	"worker/internal/redis"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Failed to load config: %s\n", err)
	}
}

func LoadRedisConfig() redis.RedisConfig {
	return redis.RedisConfig{
		Addr:        fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Username:    os.Getenv("REDIS_USERNAME"),
		Password:    os.Getenv("REDIS_PASSWORD"),
		DB:          0,
		MaxRetries:  5,
		DialTimeout: 5 * time.Second,
		Timeout:     3 * time.Second,
	}
}
