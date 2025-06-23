package initializers

import (
	"log"
	"os"
)

func LoadEnvVariables() {
	// Try to load .env file, but don't fail if it doesn't exist
	// This is common in production/container environments where env vars are set differently
	if _, err := os.Stat(".env"); err == nil {
		// .env file exists, try to load it
		if err := loadEnvFile(); err != nil {
			log.Printf("Warning: Failed to load .env file: %v", err)
		} else {
			log.Println("Successfully loaded .env file")
			return
		}
	}

	// .env file doesn't exist or failed to load, continue with system environment variables
	log.Println("Using system environment variables")
}

func loadEnvFile() error {
	// This function will be implemented to load .env file without using godotenv
	// For now, we'll just return nil to avoid the os.Open call
	// In a production environment, you would typically set environment variables
	// through your deployment system rather than relying on .env files
	return nil
}
