package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB            DBConfig
	ServerAddress string
	TokenSecret   string
}

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			Username: getEnv("DB_USERNAME", "root"),
			Password: getEnv("DB_PASSWORD", "root"),
			DbName:   getEnv("DB_NAME", "simple_bank"),
		},
		ServerAddress: getEnv("SERVER_ADDRESS", "0.0.0.0:8080"),
		TokenSecret:   getEnv("TOKEN_SECRET", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
