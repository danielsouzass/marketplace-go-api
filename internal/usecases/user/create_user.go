package user

import (
	"marketplace/internal/types"
	"net/http"
)

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Tel      string `json:"tel" validate:"required,e164"`
	Password string `json:"password" validate:"required,min=8"`
	Avatar   string `json:"avatar" validate:"url"`
}

type CreateUserResponse struct {
	Id string `json:"id"`
}

func CreateUser400Response(body types.Error) types.Response {
	return types.Response{
		Body: body,
		Code: http.StatusBadRequest,
	}
}

func CreateUser201Response(body CreateUserResponse) types.Response {
	return types.Response{
		Body: body,
		Code: http.StatusCreated,
	}
}
