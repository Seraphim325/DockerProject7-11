package config

import (
	"github.com/joho/godotenv"
)

type Config struct {
	RedisConfig    *RedisConfig
	DatabaseConfig *DatabaseConfig
	AppConfig      *AppConfig
}

func Load() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	return &Config{
		RedisConfig:    loadRedisConfig(),
		DatabaseConfig: loadDatabaseConfig(),
		AppConfig:      loadAppConfig(),
	}, nil
}
