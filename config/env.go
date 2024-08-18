package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Env is a global variable that holds the enviroment variables
var Env = initEnviroment()

type EnvType struct {
	SERVER_PORT string
	SERVER_HOST string
}

func initEnviroment() *EnvType {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
	return &EnvType{
		SERVER_PORT: getEnv("SERVER_PORT", "3000"),
		SERVER_HOST: getEnv("SERVER_HOST", "localhost"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
