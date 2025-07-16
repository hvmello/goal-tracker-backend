package user

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hvmello/goal-tracker-backend/pkg/response"
	"golang.org/x/crypto/bcrypt"
)

// getEnv is a helper to get env vars with default
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

type Service interface {
	Register(name, email, password string) (*User, error)
	Login(email, password string) (*User, string, error)
}

type service struct {
	repo      Repository
	jwtSecret []byte
}

// NewService creates a new user service
func NewService(r Repository) Service {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		// Log a warning instead of panicking
		// You might want to add proper logging here
		println("WARNING: JWT_SECRET environment variable is not set")
	}

	return &service{
		repo:      r,
		jwtSecret: []byte(jwtSecret),
	}
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

func (s *service) Login(email, password string) (*User, string, error) {
	if len(s.jwtSecret) == 0 {
		return nil, "", errors.New("JWT_SECRET environment variable is required")
	}

	user, err := s.repo.FindByEmail(email)
	if err != nil || user == nil {
		return nil, "", response.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", response.ErrInvalidCredentials
	}

	token, err := s.generateJWT(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *service) generateJWT(user *User) (string, error) {
	if len(s.jwtSecret) == 0 {
		return "", errors.New("JWT_SECRET environment variable is required")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}
