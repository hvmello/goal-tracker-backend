package integration

import (
	"github.com/hvmello/goal-tracker-backend/internal/config"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip tests")
	}

	cfg := config.GetConfig()

	db, err := config.NewDBConnection(cfg)
	if err != nil {
		t.Fatalf("Error connecting to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Error getting underlying database connection: %v", err)
	}
	defer sqlDB.Close()

	err = sqlDB.Ping()
	if err != nil {
		t.Fatalf("Error to ping database: %v", err)
	}

	t.Log("Connection successful!")
}
