package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Env is a global variable that holds the enviroment variables
var Env = initEnviroment()

type EnvType struct {
	SERVER_PORT       string
	SERVER_HOST       string
	POSTRGRES_USER    string
	POSTGRES_PASSWORD string
	POSTGRES_HOST     string
	POSTGRES_PORT     string
	POSTGRES_DBNAME   string
	JWT_SECRET        string
}

func initEnviroment() *EnvType {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
	return &EnvType{
		SERVER_PORT:       getEnv("SERVER_PORT", "3000"),
		SERVER_HOST:       getEnv("SERVER_HOST", "localhost"),
		POSTRGRES_USER:    getEnv("POSTGRES_USER", "postgres"),
		POSTGRES_PASSWORD: getEnv("POSTGRES_PASSWORD", "postgres"),
		POSTGRES_HOST:     getEnv("POSTGRES_HOST", "localhost"),
		POSTGRES_PORT:     getEnv("POSTGRES_PORT", "5432"),
		POSTGRES_DBNAME:   getEnv("POSTGRES_DBNAME", "postgres"),
		JWT_SECRET:        getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
