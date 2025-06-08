// @title Goal Tracker API
// @version 1.0
// @description API for tracking and managing goals
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://github.com/hvmello/goal-tracker-backend

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @schemes http https

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/hvmello/goal-tracker-backend/docs" // This will be auto-generated
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/hvmello/goal-tracker-backend/internal/config"
	"github.com/hvmello/goal-tracker-backend/internal/goals"
	"github.com/hvmello/goal-tracker-backend/internal/health"
)

// @title Goal Tracker API
// @version 1.0
// @description Goal tracking application API
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

	// Create router
	router := mux.NewRouter()

	// API Routes
	router.HandleFunc("/goals", goalHandler.HandleGoals).Methods("GET", "POST")
	router.HandleFunc("/goals/{id:[0-9]+}", goalHandler.HandleGoals).Methods("GET", "PUT", "DELETE")
	router.HandleFunc("/health", healthHandler.CheckHealth).Methods("GET")

	// Swagger Route
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
	))

	// Server setup
	serverAddr := fmt.Sprintf("0.0.0.0:%s", cfg.Server.Port)
	log.Printf("Server starting on port %s...", cfg.Server.Port)
	log.Printf("Swagger UI available at http://localhost:%s/swagger/index.html", cfg.Server.Port)

	if err := http.ListenAndServe(serverAddr, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// @Summary Health Check
// @Description Check if the service is healthy
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
}
