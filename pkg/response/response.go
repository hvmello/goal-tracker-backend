package response

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// WriteResponse padroniza as respostas de sucesso
func WriteResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := APIResponse{
		Success: true,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}

// WriteErrorResponse padroniza as respostas de erro
func WriteErrorResponse(w http.ResponseWriter, err error) {
	apiError, ok := err.(*APIError)
	if !ok {
		// Internal Generic Error
		apiError = &APIError{
			Message:    "Internal Server Error. Please try again later.",
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
