package user

import (
	"marketplace/internal/types"
	"net/http"
)

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func LoginUser400Response(body types.Error) types.Response {
	return types.Response{
		Body: body,
		Code: http.StatusBadRequest,
	}
}

func LoginUser201Response(body LoginResponse) types.Response {
	return types.Response{
		Body: body,
		Code: http.StatusCreated,
	}
}
