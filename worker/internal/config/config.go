package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisConfig RedisConfig
}

type RedisConfig struct {
	Addr        string
	Username    string
	Password    string
	DB          int
	MaxRetries  int
	DialTimeout time.Duration
	Timeout     time.Duration
}

func Load() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}
	return &Config{
		RedisConfig: loadRedisConfig(),
	}, nil
}

func loadRedisConfig() RedisConfig {
	return RedisConfig{
		Addr:        fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Username:    os.Getenv("REDIS_USERNAME"),
		Password:    os.Getenv("REDIS_PASSWORD"),
		DB:          0,
		MaxRetries:  5,
		DialTimeout: 5 * time.Second,
		Timeout:     3 * time.Second,
	}
}
