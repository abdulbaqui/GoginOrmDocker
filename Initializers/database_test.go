package initializers

import (
	"os"
	"testing"
)

func TestConnectToDB(t *testing.T) {
	// Set environment variables to match Docker setup
	os.Setenv("DB_HOST", "postgres")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "khan1234")
	os.Setenv("DB_NAME", "Test")
	os.Setenv("DB_PORT", "5432")

	// Skip test if not running in Docker environment
	if os.Getenv("CI") == "" && os.Getenv("DOCKER_ENV") == "" {
		t.Skip("Skipping database test - not in Docker environment")
	}

	ConnectToDB()

	// Clean up
	if DB != nil {
		sqlDB, err := DB.DB()
		if err == nil {
			sqlDB.Close()
		}
	}
}
