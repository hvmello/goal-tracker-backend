package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hvmello/goal-tracker-backend/internal/config"
	"github.com/hvmello/goal-tracker-backend/internal/goals"
)

func main() {
	db, err := config.NewDBConnection()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err := config.AutoMigrate(db); err != nil {
		log.Fatalf("Error executing migrations: %v", err)
	}

	log.Println("Migrations executed successfully!")

	// Initialize services and handlers
	goalService := goals.NewService(db)
	goalHandler := goals.NewHandler(goalService)

	// Route configuration
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"status": "ok"}
		json.NewEncoder(w).Encode(response)
	})

	http.HandleFunc("/goals/", goalHandler.HandleGoals)

	log.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
