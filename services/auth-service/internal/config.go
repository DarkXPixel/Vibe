package internal

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL         string
	JWTSecret           string
	JWTExpirationHourse int
}

func getEnv(key string, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

var config *Config

func LoadConfig() *Config {
	if config != nil {
		return config
	}

	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("No .env file found")
		godotenv.Load()
	}

	expStr := getEnv("JWT_EXPIRATION_HOURS", "24")
	exp, err := strconv.Atoi(expStr)
	if err != nil {
		log.Fatalf("Invalid JWT_EXPIRATION_HOURS: %v", err)
	}

	config = &Config{
		DatabaseURL:         getEnv("DATABASE_URL", ""),
		JWTSecret:           getEnv("JWT_SECRET", ""),
		JWTExpirationHourse: exp,
	}

	return config
}
