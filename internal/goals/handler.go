package goals

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// HandleGoals is the main handler that routes to specific methods
func (h *Handler) HandleGoals(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		h.handleGet(w, r)
	case http.MethodPost:
		h.handlePost(w, r)
	case http.MethodPut:
		h.handlePut(w, r)
	case http.MethodDelete:
		h.handleDelete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleGet processes GET requests
func (h *Handler) handleGet(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/goals/")
	if path != "" && path != "goals" {
		id, err := strconv.ParseUint(path, 10, 32)
		if err != nil {
			WriteErrorResponse(w, ErrInvalidID)
			return
		}

		goal, err := h.service.GetGoalByID(uint(id))
		if err != nil {
			WriteErrorResponse(w, ErrGoalNotFound)
			return
		}

		WriteResponse(w, http.StatusOK, goal)
		return
	}

	goals, err := h.service.GetAllGoals()
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	WriteResponse(w, http.StatusOK, goals)
}

// handlePost processes POST requests
func (h *Handler) handlePost(w http.ResponseWriter, r *http.Request) {
	var goal Goal
	if err := json.NewDecoder(r.Body).Decode(&goal); err != nil {
		WriteErrorResponse(w, ErrInvalidRequest)
		return
	}

	if err := h.service.CreateGoal(&goal); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	WriteResponse(w, http.StatusCreated, goal)
}

// handlePut processes PUT requests
func (h *Handler) handlePut(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/goals/")
	id, err := strconv.ParseUint(path, 10, 32)
	if err != nil {
		WriteErrorResponse(w, ErrInvalidID)
		return
	}

	var goal Goal
	if err := json.NewDecoder(r.Body).Decode(&goal); err != nil {
		WriteErrorResponse(w, ErrInvalidRequest)
		return
	}

	if err := h.service.UpdateGoal(uint(id), &goal); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	updatedGoal, err := h.service.GetGoalByID(uint(id))
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	WriteResponse(w, http.StatusOK, updatedGoal)
}

func (h *Handler) handleDelete(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/goals/")
	if path == "" || path == "goals" {
		WriteErrorResponse(w, ErrInvalidRequest)
		return
	}

	id, err := strconv.ParseUint(path, 10, 32)
	if err != nil {
		WriteErrorResponse(w, ErrInvalidID)
		return
	}

	if err := h.service.DeleteGoal(uint(id)); err != nil {
		WriteErrorResponse(w, err)
		return
	}

	WriteResponse(w, http.StatusNoContent, nil)
}
