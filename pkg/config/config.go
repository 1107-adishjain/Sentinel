package config

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

// Configuration settings for the Sentinel application.

type Config struct{
	PORT string 
}

func LookupEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func LoadConfig() (*Config, error) {
	// Load configuration from environment variables or default values.
	err:= godotenv.Load(".env")
	if err!=nil{
		log.Println("No .env file found, using system environment variables")
	}
	return &Config{
		PORT: LookupEnv("PORT", "8000"),
	}, nil
}

