package goals

import (
	"encoding/json"
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
)

// WriteError padronizes error message
func WriteError(w http.ResponseWriter, err error) {
	apiError, ok := err.(*APIError)
	if !ok {
		// Internal Generic Error
		apiError = &APIError{
			Message: "Internal Server Error. Please try again later. If the problem persists," +
				" contact the system administrator. Error: " + err.Error() + "",
			StatusCode: http.StatusInternalServerError,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(apiError.StatusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"error": apiError.Message,
	})
}
