package main

//go:generate go run ./cmd/terndotenv
//go:generate sqlc generate -f ./internal/store/pgstore/sqlc.yml
