package health

import (
	"github.com/hvmello/goal-tracker-backend/internal/config"
	"gorm.io/gorm"
)

type Status struct {
	Status     string            `json:"status"`
	Timestamp  string            `json:"timestamp"`
	Services   map[string]string `json:"services"`
	DbStats    map[string]int64  `json:"db_stats,omitempty"`
	AppVersion string            `json:"version"`
}

type Handler struct {
	db     *gorm.DB
	config *config.Config
}
