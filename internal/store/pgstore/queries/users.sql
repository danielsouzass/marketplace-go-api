-- name: CreateUser :one
INSERT INTO users ("name", "email", "tel", "password", "avatar")
VALUES ($1, $2, $3, $4, $5)
    RETURNING *;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: GetUserByTel :one
SELECT *
FROM users
WHERE tel = $1;