package config

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Warn("Error loading .env file, using default environment variables")
	}
	log.Info("Environment variables loaded successfully")
}

func GetEnv(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Warn("Environment variable %s not set, using default value: %s", key, defaultValue)
		return defaultValue
	}
	return value
}