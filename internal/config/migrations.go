package config

import (
	"github.com/hvmello/goal-tracker-backend/internal/goals"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&goals.Goal{})
}
