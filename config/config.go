package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds ALL environment variables for the app.
// Add new vars here — one place to maintain.
type Config struct {
	// App
	AppEnv  string
	AppPort string

	// Database
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	// JWT
	JWTAccessSecret        string
	JWTRefreshSecret       string
	JWTAccessExpiryMinutes int
	JWTRefreshExpiryDays   int
}

// AppConfig is the global singleton — initialized once at startup.
var AppConfig *Config

// Load reads the .env file (dev only) and populates AppConfig.
// Call this ONCE at the very start of main().
func Load() {
	// In production, env vars are injected by the platform (Docker, Railway, etc.)
	// godotenv.Load() silently skips if .env is missing — perfect for prod.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found — using system environment variables")
	}

	AppConfig = &Config{
		// App
		AppEnv:  getEnv("APP_ENV", "development"),
		AppPort: getEnv("APP_PORT", "8000"),

		// Database
		DBHost:     mustGetEnv("DB_HOST"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     mustGetEnv("DB_USER"),
		DBPassword: mustGetEnv("DB_PASSWORD"),
		DBName:     mustGetEnv("DB_NAME"),

		// JWT
		JWTAccessSecret:        mustGetEnv("JWT_ACCESS_SECRET"),
		JWTRefreshSecret:       mustGetEnv("JWT_REFRESH_SECRET"),
		JWTAccessExpiryMinutes: getEnvAsInt("JWT_ACCESS_EXPIRY_MINUTES", 15),
		JWTRefreshExpiryDays:   getEnvAsInt("JWT_REFRESH_EXPIRY_DAYS", 7),
	}
}

// ─── Helpers ──────────────────────────────────────────────

// getEnv returns the value or a fallback default.
func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

// mustGetEnv crashes at startup if a required variable is missing.
// Better to fail immediately than to get a runtime panic later.
func mustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("FATAL: required environment variable %q is not set", key)
	}
	return val
}

// getEnvAsInt parses an integer env var with a fallback.
func getEnvAsInt(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		if i, err := strconv.Atoi(val); err == nil {
			return i
		}
	}
	return defaultVal
}
