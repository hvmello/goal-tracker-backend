package user

import (
	"errors"

	"github.com/hvmello/goal-tracker-backend/pkg/response"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(name, email, password string) (*User, error)
	Login(email, password string) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Register(name, email, password string) (*User, error) {
	existing, _ := s.repo.FindByEmail(email)
	if existing.ID != 0 {
		return nil, errors.New("email already in use")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &User{
		Name:     name,
		Email:    email,
		Password: string(hashed),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) Login(email, password string) (*User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil || user == nil {
		return nil, response.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, response.ErrInvalidCredentials
	}
	return user, nil
}
