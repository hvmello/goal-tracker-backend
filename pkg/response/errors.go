package response

import (
	"net/http"
)

type APIError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

func (e *APIError) Error() string {
	return e.Message
}

// API Common errors
var (
	ErrInvalidID = &APIError{
		Message:    "Invalid ID",
		StatusCode: http.StatusBadRequest,
	}
	ErrGoalNotFound = &APIError{
		Message:    "Goal not found",
		StatusCode: http.StatusNotFound,
	}
	ErrInvalidRequest = &APIError{
		Message:    "Bad Request",
		StatusCode: http.StatusBadRequest,
	}
	ErrTooManyRequests = &APIError{
		Message:    "Too many requests, please try again later",
		StatusCode: http.StatusTooManyRequests,
	}
	ErrInvalidCredentials = &APIError{
		Message:    "Unauthorized access! Invalid credentials",
		StatusCode: http.StatusUnauthorized,
	}
)
