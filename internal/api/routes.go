package api

import (
	"marketplace/internal/auth"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

func (api *API) BindRoutes() {
	api.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger)

	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Post("/", api.handleCreateUser)

				r.Group(func(r chi.Router) {
					r.Use(jwtauth.Verifier(auth.TokenAuth))
					r.Use(jwtauth.Authenticator(auth.TokenAuth))
					r.Get("/me", api.handleGetUser)
				})
			})

			r.Route("/auth", func(r chi.Router) {
				r.Post("/", api.handleLoginUser)
			})
		})
	})
}
