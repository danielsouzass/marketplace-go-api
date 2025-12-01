-- name: CreateProductPaymentMethod :one
INSERT INTO product_payment_methods 
    ("product_id", "payment_method_id")
VALUES ($1, $2)
RETURNING *;

-- name: GetProductPaymentMethodsByProductID :many
SELECT pm.*
FROM product_payment_methods ppm
JOIN payment_methods pm ON ppm.payment_method_id = pm.id
WHERE ppm.product_id = $1;