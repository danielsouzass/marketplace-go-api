package api

import (
	"marketplace/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type API struct {
	Router      *chi.Mux
	Validator   *validator.Validate
	UserService services.UserService
}
