package organization

import (
	"encoding/json"
	"net/http"
	"project-management/internal/middleware"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateOrganization(w http.ResponseWriter, r *http.Request) {

	var inputData OrganizationInput
	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		http.Error(w, "name,description,image_url,banner_url is required", http.StatusBadRequest)
		return
	}
	userID := r.Context().Value(middleware.UserIDKey).(string)
	if userID == "" {
		http.Error(w, "User is not logged IN", http.StatusUnauthorized)
		return
	}
	err = h.service.CreateOrganization(r.Context(), inputData, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Organization created")
}
func (h *Handler) GetAllOrganization(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(middleware.UserIDKey).(string)
	if userID == "" {
		http.Error(w, "User is not logged IN", http.StatusUnauthorized)
		return
	}
	organizations, err := h.service.GetAllOrganizations(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(organizations)
}
func (h *Handler) GetOrganizationDetails(w http.ResponseWriter, r *http.Request) {
	orgID := chi.URLParam(r, "id")

	userID := r.Context().Value(middleware.UserIDKey).(string)
	if userID == "" {
		http.Error(w, "User is not logged IN", http.StatusUnauthorized)
		return
	}
	organization, err := h.service.GetOrganizationDetail(r.Context(), orgID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(organization)
}
func (h *Handler) DeleteOrganization(w http.ResponseWriter, r *http.Request) {
	orgID := chi.URLParam(r, "id")

	userID := r.Context().Value(middleware.UserIDKey).(string)
	if userID == "" {
		http.Error(w, "User is not logged IN", http.StatusUnauthorized)
		return
	}
	err := h.service.DeleteOrganization(r.Context(), orgID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("DELETED")
}
