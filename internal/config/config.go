package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string

	DBHost string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	return &Config{
		AppPort: getEnv("APP_PORT"),

		DBHost: getEnv("DB_HOST"),
		DBPort: getEnv("DB_PORT"),
		DBUser: getEnv("DB_USER"),
		DBPassword: getEnv("DB_PASSWORD"),
		DBName: getEnv("DB_NAME"),
	}
}

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Environment variable not set: %s", key)
	}

	return value
}
