package config

type AppConfig struct {
	Port string
}

func loadAppConfig() (*AppConfig, error) {
	return &AppConfig{
		Port: getEnv("SERVER_PORT", "8080"),
	}, nil
}
