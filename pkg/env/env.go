package env

import (
	"os"
	"strconv"
)

// db connection variables
var DatabaseUser = GetEnv("DATABASE_USER")
var DatabasePass = GetEnv("DATABASE_PASSWORD")
var DatabaseDB = GetEnv("DATABASE_DB")
var DatabaseHost = GetEnv("DATABASE_HOST")
var DatabasePort = GetEnvInt("DATABASE_PORT")

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
