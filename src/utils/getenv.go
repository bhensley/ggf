package utils

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

const DefaultTimeout = 120 * time.Second

func GetEnv(key string, defaultValue string) string {
	godotenv.Load(".env")

	// Overwrite .env with .env.production if ENV is set to production
	// Only update values that are different between environments
	if os.Getenv("ENV") == "production" {
		godotenv.Load(".env.production")
	}

	if os.Getenv(key) == "" {
		return defaultValue
	}

	return os.Getenv(key)
}
