package organization

import (
	"project-management/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func OrganizationRoutes(r chi.Router, handler *Handler) {
	r.Route("/organizations", func(r chi.Router) {
		r.Use(middleware.Auth)
		r.Post("/", handler.CreateOrganization)
		r.Get("/", handler.GetAllOrganization)
		r.Get("/{id}", handler.GetOrganizationDetails)
		r.Delete("/{id}", handler.DeleteOrganization)
	})
}
