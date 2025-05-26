package goals

import (
	"errors"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

// GetAllGoals retrieves all goals from the database
func (s *Service) GetAllGoals() ([]Goal, error) {
	var goals []Goal
	err := s.db.Find(&goals).Error
	return goals, err
}

// GetGoalByID retrieves a specific goal by its ID
func (s *Service) GetGoalByID(id uint) (*Goal, error) {
	var goal Goal
	err := s.db.First(&goal, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("goal not found")
		}
		return nil, err
	}
	return &goal, nil
}

// CreateGoal creates a new goal in the database
func (s *Service) CreateGoal(goal *Goal) error {
	if goal.Title == "" {
		return errors.New("title is required")
	}

	// Set default progress if not provided
	if goal.Progress < 0 || goal.Progress > 100 {
		return errors.New("progress must be between 0 and 100")
	}

	return s.db.Create(goal).Error
}

// UpdateGoal updates an existing goal
func (s *Service) UpdateGoal(id uint, goal *Goal) error {
	if goal.Title == "" {
		return errors.New("title is required")
	}

	if goal.Progress < 0 || goal.Progress > 100 {
		return errors.New("progress must be between 0 and 100")
	}

	// Check if goal exists
	existingGoal, err := s.GetGoalByID(id)
	if err != nil {
		return errors.New("goal not found")
	}

	// Update only allowed fields
	existingGoal.Title = goal.Title
	existingGoal.Description = goal.Description
	existingGoal.Progress = goal.Progress
	existingGoal.DueDate = goal.DueDate

	return s.db.Save(existingGoal).Error
}

// DeleteGoal deletes a goal by ID
func (s *Service) DeleteGoal(id uint) error {
	result := s.db.Delete(&Goal{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("goal not found")
	}
	return nil
}
