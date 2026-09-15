package config

import "os"

type AppConfig struct {
	Port string
}

func loadAppConfig() *AppConfig {
	return &AppConfig{
		Port: os.Getenv("SERVER_PORT"),
	}
}
