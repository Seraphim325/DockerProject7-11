package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisConfig    *RedisConfig
	DatabaseConfig *DatabaseConfig
	AppConfig      *AppConfig
}

func Load() (*Config, error) {
	if len(os.Args) > 1 && os.Args[1] == ".env.example" {
		if err := godotenv.Load(".env.example"); err != nil {
			return nil, err
		}
	}

	return &Config{
		RedisConfig:    loadRedisConfig(),
		DatabaseConfig: loadDatabaseConfig(),
		AppConfig:      loadAppConfig(),
	}, nil
}
