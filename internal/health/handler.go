package health

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/hvmello/goal-tracker-backend/internal/config" // Adicionado
	"github.com/hvmello/goal-tracker-backend/internal/goals"
	"gorm.io/gorm"
)

func NewHandler(db *gorm.DB, config *config.Config) *Handler {
	return &Handler{
		db:     db,
		config: config,
	}
}

func (h *Handler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	health := Status{
		Status:     "ok",
		Timestamp:  time.Now().Format(time.RFC3339),
		Services:   make(map[string]string),
		DbStats:    make(map[string]int64),
		AppVersion: "1.0.0", // Você pode mover isso para config
	}

	h.checkDatabaseHealth(&health)
	h.addSystemMetrics(&health)

	statusCode := http.StatusOK
	if health.Status == "error" {
		statusCode = http.StatusServiceUnavailable
	}

	goals.WriteResponse(w, statusCode, health)
}

func (h *Handler) checkDatabaseHealth(health *Status) {
	sqlDB, err := h.db.DB()
	if err != nil {
		health.Status = "error"
		health.Services["database"] = fmt.Sprintf("error: %v", err)
		return
	}

	if err := sqlDB.Ping(); err != nil {
		health.Status = "error"
		health.Services["database"] = fmt.Sprintf("error: %v", err)
		return
	}

	health.Services["database"] = "ok"

	// Adiciona estatísticas do banco
	stats := sqlDB.Stats()
	health.DbStats["open_connections"] = int64(stats.OpenConnections)
	health.DbStats["in_use"] = int64(stats.InUse)
	health.DbStats["idle"] = int64(stats.Idle)
}

func (h *Handler) addSystemMetrics(health *Status) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	health.Services["memory"] = fmt.Sprintf("%.2fMB", float64(memStats.Alloc)/1024/1024)
	health.Services["goroutines"] = fmt.Sprintf("%d", runtime.NumGoroutine())
}
