package goals

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/hvmello/goal-tracker-backend/pkg/response"
)

func TestHandleGetAllGoals(t *testing.T) {
	// Criar mock service com alguns dados de teste
	mockService := NewMockService()
	now := time.Now()
	mockService.goals = map[uint]Goal{
		1: {
			ID:          1,
			Title:       "Aprender Go",
			Description: "Estudar Go por 2 horas por dia",
			Progress:    30,
			DueDate:     now.AddDate(0, 1, 0),
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		2: {
			ID:          2,
			Title:       "Fazer exercícios",
			Description: "30 minutos de exercício diário",
			Progress:    0,
			DueDate:     now.AddDate(0, 1, 0),
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}

	tests := []struct {
		name       string
		path       string
		wantStatus int
		wantLen    int
		wantID     uint
	}{
		{
			name:       "Sucesso ao buscar todos os goals",
			path:       "/goals",
			wantStatus: http.StatusOK,
			wantLen:    2,
		},
		{
			name:       "Sucesso ao buscar goal específico",
			path:       "/goals/1",
			wantStatus: http.StatusOK,
			wantID:     1,
		},
		{
			name:       "Goal não encontrado",
			path:       "/goals/999",
			wantStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test router using gorilla/mux
			router := mux.NewRouter()
			handler := NewHandler(mockService)

			// Register the handlers
			router.HandleFunc("/goals", handler.GetAllGoals).Methods("GET")
			router.HandleFunc("/goals/{id:[0-9]+}", handler.GetGoalByID).Methods("GET")

			// Create a test request
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			w := httptest.NewRecorder()

			// Serve the request through the router
			router.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("HandleGoals() status = %v, want %v", w.Code, tt.wantStatus)
			}

			if w.Code == http.StatusOK {
				var response response.APIResponse
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Errorf("Falha ao decodificar response: %v", err)
					return
				}

				if !response.Success {
					t.Errorf("Esperava response.Success = true, got false")
					return
				}

				if tt.path == "/goals" {
					// Converter data interface{} para []Goal
					goalsData, err := json.Marshal(response.Data)
					if err != nil {
						t.Errorf("Falha ao converter dados: %v", err)
						return
					}

					var goals []Goal
					if err := json.Unmarshal(goalsData, &goals); err != nil {
						t.Errorf("Falha ao decodificar goals: %v", err)
						return
					}

					if len(goals) != tt.wantLen {
						t.Errorf("Quantidade de goals = %v, want %v", len(goals), tt.wantLen)
					}
				} else {
					// Converter data interface{} para Goal
					goalData, err := json.Marshal(response.Data)
					if err != nil {
						t.Errorf("Falha ao converter dados: %v", err)
						return
					}

					var goal Goal
					if err := json.Unmarshal(goalData, &goal); err != nil {
						t.Errorf("Falha ao decodificar goal: %v", err)
						return
					}

					if tt.wantID > 0 && goal.ID != tt.wantID {
						t.Errorf("Goal ID = %v, want %v", goal.ID, tt.wantID)
					}
				}
			}
		})
	}
}
