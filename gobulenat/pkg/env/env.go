package env

import (
	"os"
	"strconv"
)

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return "No environment variable configured"
}

func GetEnvInt(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err == nil {
		return value
	}
	return value
}
