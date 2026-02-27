package user

import (
	"encoding/json"
	"net/http"
	"project-management/internal/middleware"
)

type Handler struct {
	service Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: *service,
	}
}

func (h *Handler) GetMe(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(string)

	json.NewEncoder(w).Encode(userID)

}
