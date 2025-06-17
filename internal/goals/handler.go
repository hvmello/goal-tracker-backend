package goals

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	service GoalService
}

func NewHandler(service GoalService) *Handler {
	return &Handler{service: service}
}

// @Summary Get all goals
// @Description Get all goals
// @Tags goals
// @Produce json
// @Success 200 {array} Goal
// @Failure 500 {object} APIError
// @Router /api/goals [get]
func (h *Handler) GetAllGoals(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	goals, err := h.service.GetAllGoals()
	if err != nil {
		WriteErrorResponse(w, err)
		return
	}

	WriteResponse(w, http.StatusOK, goals)
}

// @Summary Get a goal by ID
// @Description Get a goal by its ID
// @Tags goals
// @Produce json
// @Param id path int true "Goal ID"
// @Success 200 {object} Goal
// @Failure 400 {object} APIError
// @Failure 404 {object} APIError
// @Router /api/goals/{id} [get]
func (h *Handler) GetGoalByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
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
}

// @Summary Create a goal
// @Description Create a new goal
// @Tags goals
// @Accept json
// @Produce json
// @Param goal body Goal true "Goal object"
// @Success 201 {object} Goal
// @Failure 400 {object} APIError
// @Router /api/goals [post]
func (h *Handler) CreateGoal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

// @Summary Update a goal
// @Description Update an existing goal
// @Tags goals
// @Accept json
// @Produce json
// @Param id path int true "Goal ID"
// @Param goal body Goal true "Goal object"
// @Success 200 {object} Goal
// @Failure 400 {object} APIError
// @Failure 404 {object} APIError
// @Router /api/goals/{id} [put]
func (h *Handler) UpdateGoal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
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

// @Summary Delete a goal
// @Description Delete a goal by ID
// @Tags goals
// @Param id path int true "Goal ID"
// @Success 204 "No Content"
// @Failure 400 {object} APIError
// @Failure 404 {object} APIError
// @Router /api/goals/{id} [delete]
func (h *Handler) DeleteGoal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
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
