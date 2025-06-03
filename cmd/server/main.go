package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hvmello/goal-tracker-backend/internal/config"
	"github.com/hvmello/goal-tracker-backend/internal/goals"
	"github.com/hvmello/goal-tracker-backend/internal/health"
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

	// Initialize repository and service
	goalRepo := goals.NewGormRepository(db)
	goalService := goals.NewService(goalRepo)
	goalHandler := goals.NewHandler(goalService)
	healthHandler := health.NewHandler(db, cfg)

	http.HandleFunc("/goals/", goalHandler.HandleGoals)
	http.HandleFunc("/health", healthHandler.CheckHealth)

	serverAddr := fmt.Sprintf("0.0.0.0:%s", cfg.Server.Port)
	log.Printf("Server starting on port %s...", cfg.Server.Port)
	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}
