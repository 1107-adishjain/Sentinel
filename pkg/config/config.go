package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Configuration settings for the Sentinel application.

type Config struct {
	PORT          string
	RedisAddr     string
	RedisPassword string
	DBHost        string
	DBPort        string
	DBName        string
	DBUser        string
	DBPassword    string
}

func LookupEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func LoadConfig() (*Config, error) {
	// Load configuration from environment variables or default values.
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}
	return &Config{
		PORT:          LookupEnv("PORT", "8000"),
		RedisAddr:     LookupEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword: LookupEnv("REDIS_PASSWORD", ""),
		DBHost:        LookupEnv("DB_HOST", "localhost"),
		DBPort:        LookupEnv("DB_PORT", "5432"),
		DBName:        LookupEnv("DB_NAME", "sentinel"),
		DBUser:        LookupEnv("DB_USER", "postgres"),
		DBPassword:    LookupEnv("DB_PASSWORD", "password"),
	}, nil
}
