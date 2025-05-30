package goals

import "errors"

// GoalService define a interface para o serviço
type GoalService interface {
	GetAllGoals() ([]Goal, error)
	GetGoalByID(id uint) (*Goal, error)
	CreateGoal(goal *Goal) error
	UpdateGoal(id uint, goal *Goal) error
	DeleteGoal(id uint) error
}

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
	return nil, ErrGoalNotFound
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
		return ErrGoalNotFound
	}
	m.goals[id] = *goal
	return nil
}

func (m *MockService) DeleteGoal(id uint) error {
	if _, exists := m.goals[id]; !exists {
		return ErrGoalNotFound
	}
	delete(m.goals, id)
	return nil
}
