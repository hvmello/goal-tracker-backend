package goals

import (
	"github.com/hvmello/goal-tracker-backend/pkg/response"
)

// Service handles business logic for goals
type Service struct {
	repo Repository
}

// NewService creates a new instance of Service
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// GetAllGoals retrieves all goals
func (s *Service) GetAllGoals() ([]Goal, error) {
	goals, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return goals, nil
}

// GetGoalByID retrieves a specific goal by its ID
func (s *Service) GetGoalByID(id uint) (*Goal, error) {
	goal, err := s.repo.FindByID(id)
	if err != nil {
		return nil, response.ErrGoalNotFound
	}
	return goal, nil
}

// CreateGoal creates a new goal
func (s *Service) CreateGoal(goal *Goal) error {
	if err := ValidateGoal(goal); err != nil {
		return err
	}
	return s.repo.Create(goal)
}

// UpdateGoal updates an existing goal
func (s *Service) UpdateGoal(id uint, goal *Goal) error {
	if err := ValidateGoal(goal); err != nil {
		return err
	}

	existingGoal, err := s.repo.FindByID(id)
	if err != nil {
		return response.ErrGoalNotFound
	}

	// Update only allowed fields
	existingGoal.Title = goal.Title
	existingGoal.Description = goal.Description
	existingGoal.Progress = goal.Progress
	existingGoal.DueDate = goal.DueDate

	return s.repo.Update(existingGoal)
}

// DeleteGoal removes a goal
func (s *Service) DeleteGoal(id uint) error {
	if _, err := s.repo.FindByID(id); err != nil {
		return response.ErrGoalNotFound
	}

	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
