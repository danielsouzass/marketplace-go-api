-- name: CreateProduct :one
INSERT INTO products 
    ("user_id", "name", "description", "is_new", "price", "accept_trade")
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetProductByID :one
SELECT *
FROM products 
WHERE id = $1;

-- name: GetProductsByUserID :many
SELECT *
FROM products 
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: DeleteProductByID :exec
DELETE FROM products 
WHERE id = $1;
