package goals

import (
	"time"

	"gorm.io/gorm"
)

// Goal represents a user goal
// @Description A goal that a user wants to achieve
type Goal struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" gorm:"not null"`
	Description string         `json:"description"`
	DueDate     time.Time      `json:"dueDate"`
	Progress    int            `json:"progress" gorm:"default:0"`
	UserID      uint           `json:"userId"` // Foreign key
	CreatedAt   time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index" swaggerignore:"true"`
}
