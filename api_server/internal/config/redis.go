package config

import (
	"fmt"
	"os"
	"time"
)

type RedisConfig struct {
	Addr        string
	Username    string
	Password    string
	DB          int
	MaxRetries  int
	DialTimeout time.Duration
	Timeout     time.Duration
	Pattern     string
}

func loadRedisConfig() *RedisConfig {
	return &RedisConfig{
		Addr:        fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Username:    os.Getenv("REDIS_USERNAME"),
		Password:    os.Getenv("REDIS_PASSWORD"),
		DB:          0,
		MaxRetries:  5,
		DialTimeout: 5 * time.Second,
		Timeout:     3 * time.Second,
		Pattern:     "index:",
	}
}
