package config

import (
	"path/filepath"
	"log"
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	// Load environment variables
	envPath := filepath.Join("..", ".env")

	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}