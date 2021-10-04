package api

import (
	"fmt"
	"os"
)

var dbConninfo = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
	getEnv("DB_HOST", "localhost"),
	getEnv("DB_PORT", "5432"),
	getEnv("DB_USER", "postgres"),
	getEnv("DB_NAME", "postgres"),
	getEnv("DB_SSLMODE", "disable"),
	getEnv("DB_PASSWORD", "postgres"),
)

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
