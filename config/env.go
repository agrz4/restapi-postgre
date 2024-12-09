package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPORT     string
	DBUser     string
	DBPassword string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		DBHost:     getENV("DB_HOST", "127.0.0.1"),
		DBPORT:     getENV("DB_PORT", "5432"),
		DBUser:     getENV("DB_USER", "youruser"),
		DBPassword: getENV("DB_PASSWORD", "yourpassword"),
		DBName:     getENV("DB_NAME", "yourdbname"),
	}
}

func getENV(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
