package integration

import (
	"github.com/hvmello/goal-tracker-backend/internal/config"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip tests")
	}

	db, err := config.NewDBConnection()
	if err != nil {
		t.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Fatalf("Error to ping database: %v", err)
	}

	t.Log("Connection successful!")
}
