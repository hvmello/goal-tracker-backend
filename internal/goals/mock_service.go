package goals

import (
	"errors"

	"github.com/hvmello/goal-tracker-backend/pkg/response"
)

// MockService implementa GoalService para testes
type MockService struct {
	goals map[uint]Goal
}

func NewMockService() *MockService {
	return &MockService{
		goals: make(map[uint]Goal),
	}
}

func (m *MockService) GetAllGoals() ([]Goal, error) {
	goals := make([]Goal, 0, len(m.goals))
	for _, goal := range m.goals {
		goals = append(goals, goal)
	}
	return goals, nil
}

func (m *MockService) GetGoalByID(id uint) (*Goal, error) {
	if goal, exists := m.goals[id]; exists {
		return &goal, nil
	}
	return nil, response.ErrGoalNotFound
}

func (m *MockService) CreateGoal(goal *Goal) error {
	if goal.ID == 0 {
		return errors.New("ID can't be zero")
	}
	m.goals[goal.ID] = *goal
	return nil
}

func (m *MockService) UpdateGoal(id uint, goal *Goal) error {
	if _, exists := m.goals[id]; !exists {
		return response.ErrGoalNotFound
	}
	m.goals[id] = *goal
	return nil
}

func (m *MockService) DeleteGoal(id uint) error {
	if _, exists := m.goals[id]; !exists {
		return response.ErrGoalNotFound
	}
	delete(m.goals, id)
	return nil
}

func (m *MockService) PartialUpdateGoal(id uint, req *UpdateGoalRequest) (*Goal, error) {
	goal, exists := m.goals[id]
	if !exists {
		return nil, response.ErrGoalNotFound
	}

	// Update only the fields that are provided in the request
	if req.Title != nil {
		goal.Title = *req.Title
	}
	if req.Description != nil {
		goal.Description = *req.Description
	}
	if req.Progress != nil {
		goal.Progress = *req.Progress
	}
	if req.DueDate != nil {
		goal.DueDate = *req.DueDate
	}

	m.goals[id] = goal

	return nil, nil
}
