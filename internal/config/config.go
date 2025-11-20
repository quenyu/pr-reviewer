package config

import "os"

const (
	defaultHTTPPort    = "8080"
	defaultDatabaseDSN = "postgres://pr-reviewer:secret@localhost:5432/pr-reviewer?sslmode=disable"
)

type Config struct {
	HTTPPort    string
	DatabaseDSN string
}

func Load() Config {
	cfg := Config{
		HTTPPort:    getenv("PORT", defaultHTTPPort),
		DatabaseDSN: getenv("DATABASE_URL", defaultDatabaseDSN),
	}

	return cfg
}

func getenv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return fallback
}
