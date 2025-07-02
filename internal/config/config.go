package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Server    ServerConfig
	Database  DatabaseConfig
	RateLimit RateLimitConfig
}

type ServerConfig struct {
	Port    string // Port is the port on which the server will listen
	BaseURL string // BaseURL is optional and can be used to set a base path for the server
}

type DatabaseConfig struct {
	Host     string // Host is the database host, e.g., "localhost" or an IP address
	Port     string // Port is the database port, e.g., "5432" for PostgreSQL
	User     string
	Password string
	DBName   string
	SSLMode  string // SSLMode is the SSL mode for the database connection, e.g., "disable", "require", etc.
}

type RateLimitConfig struct {
	Enabled           bool // Enabled indicates whether rate limiting is enabled
	RequestsPerMinute int
	BurstSize         int
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
			log.Printf(".env file loaded from: %s", path)
			loadedEnv = true
			break
		}
	}

	if !loadedEnv {
		log.Printf("Warning: .env file not found in any of the search paths")
	}

	cfg := &Config{
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
			RequestsPerMinute: getEnvInt("RATE_LIMIT_REQUESTS_PER_MINUTE", 60),
			BurstSize:         getEnvInt("RATE_LIMIT_BURST_SIZE", 20),
		},
	}

	if cfg.Database.Host == "" || cfg.Database.User == "" || cfg.Database.Password == "" || cfg.Database.DBName == "" {
		log.Fatal("Database infos are incomplete. Check the env variables.")
	}

	if cfg.Database.Host == "" || cfg.Database.User == "" || cfg.Database.Password == "" || cfg.Database.DBName == "" {
		log.Fatal("Database infos are incomplete. Check the env variables: DB_HOST, DB_USER, DB_PASSWORD, DB_NAME.")
	}
	if cfg.RateLimit.RequestsPerMinute <= 0 {
		log.Fatal("Rate limit RequestsPerMinute must be greater than zero.")
	}
	if cfg.RateLimit.BurstSize <= 0 {
		log.Fatal("Rate limit BurstSize must be greater than zero.")
	}

	return cfg
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
