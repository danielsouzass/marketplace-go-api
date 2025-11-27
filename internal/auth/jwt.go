package auth

import (
	"os"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var TokenAuth *jwtauth.JWTAuth

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	TokenAuth = jwtauth.New("HS256", []byte(os.Getenv("MARKETPLACE_JWT_SECRET")), nil)
}

func NewAccessToken(userId uuid.UUID) (accessToken string, err error) {
	_, tokenString, err := TokenAuth.Encode(map[string]any{
		"user_id": userId.String(),
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	})
	return tokenString, err
}
