-- name: CreateProduct :one
INSERT INTO products 
    ("user_id", "name", "description", "is_new", "price", "accept_trade")
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;