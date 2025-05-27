// Novo arquivo: internal/goals/response.go
package goals

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func WriteResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := APIResponse{
		Success: true,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}

func WriteErrorResponse(w http.ResponseWriter, err error) {
	apiError, ok := err.(*APIError)
	if !ok {
		apiError = &APIError{
			Message:    "Erro interno do servidor",
			StatusCode: http.StatusInternalServerError,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(apiError.StatusCode)

	response := APIResponse{
		Success: false,
		Error:   apiError.Message,
	}

	json.NewEncoder(w).Encode(response)
}
