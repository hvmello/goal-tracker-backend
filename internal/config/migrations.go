package config

import (
	"github.com/hvmello/goal-tracker-backend/internal/goals"
	"github.com/hvmello/goal-tracker-backend/internal/user"
	"gorm.io/gorm"
)

// AutoMigrate executes database migrations for all models
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&user.User{},
		&goals.Goal{},
	)
}
