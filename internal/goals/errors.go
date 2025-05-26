package goals

import (
	"encoding/json"
	"net/http"
)

// APIError representa um erro da API com status HTTP
type APIError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"-"` // não será serializado
}

func (e *APIError) Error() string {
	return e.Message
}

// Erros comuns da API
var (
	ErrInvalidID = &APIError{
		Message:    "ID inválido",
		StatusCode: http.StatusBadRequest,
	}
	ErrGoalNotFound = &APIError{
		Message:    "Meta não encontrada",
		StatusCode: http.StatusNotFound,
	}
	ErrInvalidRequest = &APIError{
		Message:    "Requisição inválida",
		StatusCode: http.StatusBadRequest,
	}
)

// WriteError escreve o erro na resposta HTTP de forma padronizada
func WriteError(w http.ResponseWriter, err error) {
	apiError, ok := err.(*APIError)
	if !ok {
		// Erro interno genérico
		apiError = &APIError{
			Message:    "Erro interno do servidor",
			StatusCode: http.StatusInternalServerError,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(apiError.StatusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"error": apiError.Message,
	})
}
