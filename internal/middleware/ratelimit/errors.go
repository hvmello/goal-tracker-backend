package ratelimit

import (
	"github.com/hvmello/goal-tracker-backend/internal/goals"
	"net/http"
)

var (
	ErrTooManyRequests = &goals.APIError{
		Message:    "Too many requests. Please try again later.",
		StatusCode: http.StatusTooManyRequests,
	}
)
