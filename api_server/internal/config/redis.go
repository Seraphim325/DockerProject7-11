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
	Sentinel    string
}

func loadRedisConfig() (*RedisConfig, error) {

	host, err := requireEnv("REDIS_HOST")
	if err != nil {
		return nil, err
	}

	user, err := requireEnv("REDIS_USERNAME")
	if err != nil {
		return nil, err
	}

	password, err := getSecret("REDIS_PASSWORD")

	return &RedisConfig{
		Addr:        fmt.Sprintf("%s:%s", host, os.Getenv("REDIS_PORT")),
		Username:    user,
		Password:    password,
		DB:          0,
		MaxRetries:  5,
		DialTimeout: 5 * time.Second,
		Timeout:     3 * time.Second,
		Pattern:     "index:",
		Sentinel:    "__NULL__",
	}, nil
}
