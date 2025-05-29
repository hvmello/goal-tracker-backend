package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hvmello/goal-tracker-backend/internal/config"
	"github.com/hvmello/goal-tracker-backend/internal/goals"
)

func main() {
	cfg := config.GetConfig()

	db, err := config.NewDBConnection(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err := config.AutoMigrate(db); err != nil {
		log.Fatalf("Error executing migrations: %v", err)
	}

	goalService := goals.NewService(db)
	goalHandler := goals.NewHandler(goalService)

	http.HandleFunc("/goals/", goalHandler.HandleGoals)
	http.HandleFunc("/health", healthCheck)

	serverAddr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Server starting on port %s...", cfg.Server.Port)
	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}
