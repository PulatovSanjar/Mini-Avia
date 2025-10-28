package common

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port        int
	LogLevel    string
	DatabaseURL string
}

func MustLoad() Config {
	port := 8080
	if v := os.Getenv("APP_PORT"); v != "" {
		if p, err := strconv.Atoi(v); err == nil {
			port = p
		}
	}
	cfg := Config{
		Port:        port,
		LogLevel:    getenv("LOG_LEVEL", "info"),
		DatabaseURL: mustEnv("DATABASE_URL"),
	}
	return cfg
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func mustEnv(key string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	panic(fmt.Sprintf("environment variable %s not set", key))
}
