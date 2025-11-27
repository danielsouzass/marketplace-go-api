package user

import (
	"marketplace/internal/types"
	"net/http"
	"time"
)

type GetUserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Tel       string    `json:"tel"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetUser404Response(body types.Error) types.Response {
	return types.Response{
		Body: body,
		Code: http.StatusNotFound,
	}
}

func GetUser200Response(body GetUserResponse) types.Response {
	return types.Response{
		Body: body,
		Code: http.StatusOK,
	}
}
