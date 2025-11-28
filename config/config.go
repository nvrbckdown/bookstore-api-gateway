package config

import (
	"os"
)

type Config struct {
	Port                string
	GinMode             string
	BookstoreServiceURL string
	OrderServiceURL     string
	JWTSecret           string
	TokenExpiry         string
	AllowedOrigins      string
}

func LoadConfig() *Config {
	return &Config{
		Port:                getEnv("PORT", "8080"),
		GinMode:             getEnv("GIN_MODE", "debug"),
		BookstoreServiceURL: getEnv("BOOKSTORE_SERVICE_URL", "http://localhost:5000"),
		OrderServiceURL:     getEnv("ORDER_SERVICE_URL", "http://localhost:8000"),
		JWTSecret:           getEnv("JWT_SECRET", "mysecret"),
		TokenExpiry:         getEnv("TOKEN_EXPIRY", "24h"),
		AllowedOrigins:      getEnv("ALLOWED_ORIGINS", "*"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
