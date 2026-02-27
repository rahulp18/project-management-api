package user

import (
	"project-management/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(r chi.Router, handler *Handler) {
	r.Route("/user", func(r chi.Router) {
		r.Use(middleware.Auth)

		r.Get("/me", handler.GetMe)
	})
}
