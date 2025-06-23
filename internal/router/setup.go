package router

import (
	"log"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/hvmello/goal-tracker-backend/internal/config"
	"github.com/hvmello/goal-tracker-backend/internal/goals"
	"github.com/hvmello/goal-tracker-backend/internal/health"
	"github.com/hvmello/goal-tracker-backend/internal/middleware/ratelimit"
	"github.com/hvmello/goal-tracker-backend/internal/middleware/securityheaders"
	"github.com/hvmello/goal-tracker-backend/internal/user"
)

func SetupRouter(cfg *config.Config) *mux.Router {
	r := mux.NewRouter()

	// Security headers middleware
	r.Use(securityheaders.Middleware)

	// Rate limiting middleware
	rateLimiter := ratelimit.New(cfg.RateLimit)
	r.Use(rateLimiter.Middleware)

	// Swagger
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
	))

	// Database
	db, err := config.NewDBConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}

	if err := config.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Health
	healthHandler := health.NewHandler(db, cfg)
	r.HandleFunc("/health", healthHandler.CheckHealth).Methods("GET")

	// Goals
	goalRepo := goals.NewGormRepository(db)
	goalService := goals.NewService(goalRepo)
	goalHandler := goals.NewHandler(goalService)

	r.HandleFunc("/api/goals", goalHandler.GetAllGoals).Methods("GET")
	r.HandleFunc("/api/goals/{id:[0-9]+}", goalHandler.GetGoalByID).Methods("GET")
	r.HandleFunc("/api/goals", goalHandler.CreateGoal).Methods("POST")
	r.HandleFunc("/api/goals/{id:[0-9]+}", goalHandler.UpdateGoal).Methods("PUT")
	r.HandleFunc("/api/goals/{id:[0-9]+}", goalHandler.DeleteGoal).Methods("DELETE")

	// Users
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	r.HandleFunc("/api/users/register", userHandler.Register).Methods("POST")
	// Adicione outras rotas de usuário conforme necessário

	return r
}
