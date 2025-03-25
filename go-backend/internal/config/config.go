package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config contains application configuration
type Config struct {
	Port             string
	DatabaseURL      string
	JWTSecret        string
	JWTExpiryHours   int
	Environment      string
	AllowedOrigins   []string
	TranscriptAPIURL string
	SentimentAPIURL  string
}

// Load loads the configuration from environment variables
func Load() (*Config, error) {
	// Load .env file if exists
	godotenv.Load()

	// Set defaults and get overrides from env vars
	config := &Config{
		Port:             getEnv("PORT", "8080"),
		DatabaseURL:      getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/callguard?sslmode=disable"),
		JWTSecret:        getEnv("JWT_SECRET", "secret-key"),
		Environment:      getEnv("ENVIRONMENT", "development"),
		TranscriptAPIURL: getEnv("TRANSCRIPT_API_URL", ""),
		SentimentAPIURL:  getEnv("SENTIMENT_API_URL", ""),
	}

	// Parse allowed origins
	allowedOriginsStr := getEnv("ALLOWED_ORIGINS", "http://localhost:3000")
	if allowedOriginsStr != "" {
		config.AllowedOrigins = []string{allowedOriginsStr}
	}

	// Parse integer values
	var err error
	config.JWTExpiryHours, err = strconv.Atoi(getEnv("JWT_EXPIRY_HOURS", "24"))
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_EXPIRY_HOURS: %w", err)
	}

	return config, nil
}

// getEnv retrieves an environment variable or returns the fallback
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
