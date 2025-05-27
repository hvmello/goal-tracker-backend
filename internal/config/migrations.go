package config

import (
	"github.com/hvmello/goal-tracker-backend/internal/goals"
	"gorm.io/gorm"
)

// AutoMigrate executes database migrations for all models
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&goals.Goal{},
	)
}
