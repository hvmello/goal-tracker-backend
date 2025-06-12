package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type Config struct {
	Server    ServerConfig
	Database  DatabaseConfig
	RateLimit RateLimitConfig
}

type ServerConfig struct {
	Port    string
	BaseURL string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type RateLimitConfig struct {
	Enabled           bool `env:"RATE_LIMIT_ENABLED" envDefault:"true"`
	RequestsPerMinute int  `env:"RATE_LIMIT_REQUESTS_PER_MINUTE" envDefault:"60"`
	BurstSize         int  `env:"RATE_LIMIT_BURST_SIZE" envDefault:"20"`
}

func GetConfig() *Config {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("Warning: Could not get current directory: %v", err)
	}

	// Try to load .env from the current directory and parent directory
	envPaths := []string{
		".env",
		filepath.Join(currentDir, ".env"),
		filepath.Join(filepath.Dir(currentDir), ".env"),
	}

	var loadedEnv bool
	for _, path := range envPaths {
		if err := godotenv.Load(path); err == nil {
			loadedEnv = true
			break
		}
	}

	if !loadedEnv {
		log.Printf("Warning: .env file not found in any of the search paths")
	}

	return &Config{
		Server: ServerConfig{
			Port:    getEnv("SERVER_PORT", "8080"),
			BaseURL: getEnv("SERVER_BASE_URL", ""),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "goaltracker"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		RateLimit: RateLimitConfig{
			Enabled:           getEnvBool("RATE_LIMIT_ENABLED", true),
			RequestsPerMinute: getEnvInt("RATE_LIMIT_REQUESTS_PER_MINUTE", 5),
			BurstSize:         getEnvInt("RATE_LIMIT_BURST_SIZE", 20),
		},
	}
}

// getEnv retrieves an environment variable or returns the default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetEnvBool retrieves a boolean environment variable
func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		b, err := strconv.ParseBool(value)
		if err == nil {
			return b
		}
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
	}
	return defaultValue
}
