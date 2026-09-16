package config

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func loadDatabaseConfig() (*DatabaseConfig, error) {

	host, err := requireEnv("DB_HOST")
	if err != nil {
		return nil, err
	}

	user, err := requireEnv("DB_USER")
	if err != nil {
		return nil, err
	}

	password, err := getSecret("DB_PASSWORD")
	if err != nil {
		return nil, err
	}

	name, err := requireEnv("DB_NAME")
	if err != nil {
		return nil, err
	}

	return &DatabaseConfig{
		Host:     host,
		Port:     getEnv("DB_PORT", "5432"),
		User:     user,
		Password: password,
		Name:     name,
		SSLMode:  getEnv("DB_SSLMODE", "disabled"),
	}, nil
}
