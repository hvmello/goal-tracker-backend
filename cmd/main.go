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
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/hvmello/goal-tracker-backend/docs" // This will be auto-generated
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/hvmello/goal-tracker-backend/internal/config"
	"github.com/hvmello/goal-tracker-backend/internal/goals"
	"github.com/hvmello/goal-tracker-backend/internal/health"
	"github.com/hvmello/goal-tracker-backend/internal/middleware/ratelimit"
	"github.com/hvmello/goal-tracker-backend/internal/router"
)

// @title Goal Tracker API
// @version 1.0
// @description Goal tracking application API
func main() {
	cfg := loadConfiguration()
	r := setupServer(cfg)
	server := createHTTPServer(cfg, r)

	startServer(server)
	waitForShutdown(server)
}

func loadConfiguration() *config.Config {
	return config.GetConfig()
}

func setupServer(cfg *config.Config) *router.Router {
	r := router.NewRouter()

	setupMiddlewares(r, cfg)
	setupRoutes(r)

	return r
}

func setupMiddlewares(r *router.Router, cfg *config.Config) {
	log.Printf("Setting up middlewares. Rate limit enabled: %v", cfg.RateLimit.Enabled)

	rateLimiter := ratelimit.New(cfg.RateLimit)
	r.Use(rateLimiter.Middleware)
}

func setupRoutes(r *router.Router) {
	// Database connection
	cfg := loadConfiguration()
	db, err := config.NewDBConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}

	// Run migrations
	if err := config.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Health endpoint
	healthHandler := health.NewHandler(db, cfg)

	// Goals endpoints
	goalRepo := goals.NewGormRepository(db)
	goalService := goals.NewService(goalRepo)
	goalHandler := goals.NewHandler(goalService)

	r.Handle("/health", http.HandlerFunc(healthHandler.CheckHealth))
	r.Handle("/api/goals", http.HandlerFunc(goalHandler.HandleGoals))
	r.Handle("/api/goals/", http.HandlerFunc(goalHandler.HandleGoals))

	// Swagger Route
	r.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
	))

}

func createHTTPServer(cfg *config.Config, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.Server.Port),
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

func startServer(server *http.Server) {
	go func() {
		log.Printf("Server starting on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()
}

func waitForShutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped gracefully")
}
