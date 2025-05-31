package goals

import (
	"gorm.io/gorm"
)

// Repository defines the interface for data access methods
type Repository interface {
	FindAll() ([]Goal, error)
	FindByID(id uint) (*Goal, error)
	Create(goal *Goal) error
	Update(goal *Goal) error
	Delete(id uint) error
}

// GormRepository implements Repository using GORM
type GormRepository struct {
	db *gorm.DB
}

// NewGormRepository creates a new instance of GormRepository
func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

// FindAll retrieves all goals from the database
func (r *GormRepository) FindAll() ([]Goal, error) {
	var goals []Goal
	err := r.db.Find(&goals).Error
	return goals, err
}

// FindByID retrieves a specific goal by its ID
func (r *GormRepository) FindByID(id uint) (*Goal, error) {
	var goal Goal
	err := r.db.First(&goal, id).Error
	if err != nil {
		return nil, err
	}
	return &goal, nil
}

// Create adds a new goal to the database
func (r *GormRepository) Create(goal *Goal) error {
	return r.db.Create(goal).Error
}

// Update modifies an existing goal in the database
func (r *GormRepository) Update(goal *Goal) error {
	return r.db.Save(goal).Error
}

// Delete removes a goal from the database
func (r *GormRepository) Delete(id uint) error {
	return r.db.Delete(&Goal{}, id).Error
}
