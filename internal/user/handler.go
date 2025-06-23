package user

import (
	"encoding/json"
	"net/http"

	"github.com/hvmello/goal-tracker-backend/pkg/response"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteErrorResponse(w, err)
		return
	}

	user, err := h.service.Register(req.Name, req.Email, req.Password)
	if err != nil {
		response.WriteErrorResponse(w, err)
		return
	}

	resp := map[string]interface{}{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	}
	response.WriteResponse(w, http.StatusCreated, resp)
}
