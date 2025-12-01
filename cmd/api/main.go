package main

import (
	"context"
	"fmt"
	"marketplace/internal/api"
	"marketplace/internal/services"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("MARKETPLACE_DATABASE_USER"),
		os.Getenv("MARKETPLACE_DATABASE_PASSWORD"),
		os.Getenv("MARKETPLACE_DATABASE_HOST"),
		os.Getenv("MARKETPLACE_DATABASE_PORT"),
		os.Getenv("MARKETPLACE_DATABASE_NAME"),
	))

	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	api := api.API{
		Router:         chi.NewMux(),
		Validator:      validator.New(validator.WithRequiredStructEnabled()),
		UserService:    services.NewUserService(pool),
		ProductService: services.NewProductService(pool),
	}

	api.BindRoutes()

	fmt.Println("Starting server on port: 3000")
	if err := http.ListenAndServe("localhost:3000", api.Router); err != nil {
		panic(err)
	}
}
