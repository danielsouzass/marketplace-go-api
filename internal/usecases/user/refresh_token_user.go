package user

import (
	"marketplace/internal/types"
	"net/http"
)

type RefreshTokenUserRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshTokenUserResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func RefreshTokenUser400Response(body types.Error) types.Response {
	return types.Response{
		Body: body,
		Code: http.StatusBadRequest,
	}
}

func RefreshTokenUser200Response(body RefreshTokenUserResponse) types.Response {
	return types.Response{
		Body: body,
		Code: http.StatusOK,
	}
}
