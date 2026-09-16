package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppConfig      *AppConfig
	RedisConfig    *RedisConfig
	DatabaseConfig *DatabaseConfig
}

func Load() (*Config, error) {
	if len(os.Args) > 1 && os.Args[1] == ".env.example" {
		if err := godotenv.Load(".env.example"); err != nil {
			return nil, err
		}
	}

	appConf, err := loadAppConfig()
	if err != nil {
		return nil, err
	}

	redisConf, err := loadRedisConfig()
	if err != nil {
		return nil, err
	}

	dbConf, err := loadDatabaseConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		AppConfig:      appConf,
		RedisConfig:    redisConf,
		DatabaseConfig: dbConf,
	}, nil
}

func getSecret(name string) (string, error) {
	if v := os.Getenv(name); v != "" {
		return v, nil
	}

	if path := os.Getenv(name + "_FILE"); path != "" {
		content, err := os.ReadFile(path)
		if err != nil {
			return "", fmt.Errorf("error reading secret %s from %s: %w\n", name, path, err)
		}
		return strings.TrimSpace(string(content)), nil
	}

	return "", fmt.Errorf("secret file %s is not specified\n", name)
}

func getEnv(key string, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return defaultValue
}

func requireEnv(key string) (string, error) {
	if v := os.Getenv(key); v != "" {
		return v, nil
	}
	return "", fmt.Errorf("required variable %s is not specified\n", key)
}
