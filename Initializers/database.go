package initializers

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// getEnvWithDefault safely gets environment variable with a default value
func getEnvWithDefault(key, defaultValue string) string {
	// Try to get from environment, but don't fail if it's not available
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// ConnectToDB connects to the database with fallback defaults
func ConnectToDB() error {
	// Use default values to avoid file system operations
	host := getEnvWithDefault("DB_HOST", "localhost")
	user := getEnvWithDefault("DB_USER", "postgres")
	password := getEnvWithDefault("DB_PASSWORD", "postgres")
	dbname := getEnvWithDefault("DB_NAME", "postgres")
	port := getEnvWithDefault("DB_PORT", "5432")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	log.Printf("Attempting to connect to database with DSN: %s", dsn)

	// Use minimal GORM config to avoid file system operations
	config := &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent), // Disable SQL logging to avoid file operations
		DisableForeignKeyConstraintWhenMigrating: true,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}

	log.Println("Successfully connected to database")
	return nil
}

// ConnectToDBContainer connects to database using container defaults (no file system ops)
func ConnectToDBContainer() error {
	// Use hardcoded defaults for container environments
	dsn := "host=postgres user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"

	log.Printf("Attempting to connect to database with container DSN: %s", dsn)

	// Use minimal GORM config to avoid file system operations
	config := &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent), // Disable SQL logging to avoid file operations
		DisableForeignKeyConstraintWhenMigrating: true,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}

	log.Println("Successfully connected to database")
	return nil
}

// ConnectToDBMinimal connects to database with absolute minimal configuration
func ConnectToDBMinimal() error {
	// Use hardcoded defaults - no environment variable access
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"

	log.Printf("Attempting to connect to database with minimal DSN: %s", dsn)

	// Use absolute minimal GORM config - no file system operations
	config := &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Error), // Only log errors
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true, // Avoid transaction logging
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}

	log.Println("Successfully connected to database")
	return nil
}

// ConnectToDBDirect connects to database with direct driver configuration to avoid os.Stat
func ConnectToDBDirect() error {
	// Use hardcoded connection string with minimal configuration
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=UTC"

	log.Printf("Attempting to connect to database with direct DSN: %s", dsn)

	// Use the most minimal configuration possible
	config := &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Error), // Only log errors
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		DisableAutomaticPing:                     true, // Avoid ping operations that might use file system
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}

	log.Println("Successfully connected to database with direct connection")
	return nil
}
