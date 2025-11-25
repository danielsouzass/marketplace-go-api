-- name: CreateUser :one
INSERT INTO users ("name", "email", "tel", "password", "avatar")
VALUES ($1, $2, $3, $4, $5)
    RETURNING *;