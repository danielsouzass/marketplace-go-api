-- name: CreateProductImage :one
INSERT INTO product_images 
    ("product_id", "path")
VALUES ($1, $2)
RETURNING *;