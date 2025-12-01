-- name: CreateProductPaymentMethod :one
INSERT INTO product_payment_methods 
    ("product_id", "payment_method_id")
VALUES ($1, $2)
RETURNING *;