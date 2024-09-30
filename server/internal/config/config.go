// internal/config/config.go
package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseConfig *DatabaseConfig
	ServerPort     string
	LogLevel       string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func Load() (*Config, error) {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	databaseConfig := &DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}

	config := &Config{
		DatabaseConfig: databaseConfig,
		ServerPort:     os.Getenv("SERVER_PORT"),
		LogLevel:       os.Getenv("LOG_LEVEL"),
	}

	return config, nil
}
