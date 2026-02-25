package auth

import "github.com/go-chi/chi/v5"

func AuthRoutes(r chi.Router, handler *Handler) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", handler.Register)
	})
}
