package goals

import (
	"github.com/hvmello/goal-tracker-backend/pkg/response"
)

// MockRepository implements Repository interface for testing
type MockRepository struct {
	goals map[uint]Goal
}

// NewMockRepository creates a new instance of MockRepository
func NewMockRepository() *MockRepository {
	return &MockRepository{
		goals: make(map[uint]Goal),
	}
}

func (m *MockRepository) FindAll() ([]Goal, error) {
	goals := make([]Goal, 0, len(m.goals))
	for _, goal := range m.goals {
		goals = append(goals, goal)
	}
	return goals, nil
}

func (m *MockRepository) FindByID(id uint) (*Goal, error) {
	if goal, exists := m.goals[id]; exists {
		return &goal, nil
	}
	return nil, response.ErrGoalNotFound
}

func (m *MockRepository) Create(goal *Goal) error {
	if goal.ID == 0 {
		return response.ErrInvalidID
	}
	m.goals[goal.ID] = *goal
	return nil
}

func (m *MockRepository) Update(goal *Goal) error {
	if _, exists := m.goals[goal.ID]; !exists {
		return response.ErrGoalNotFound
	}
	m.goals[goal.ID] = *goal
	return nil
}

func (m *MockRepository) Delete(id uint) error {
	if _, exists := m.goals[id]; !exists {
		return response.ErrGoalNotFound
	}
	delete(m.goals, id)
	return nil
}
