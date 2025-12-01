-- name: GetPaymentMethodsByKeys :many
SELECT id, key, name
FROM payment_methods
WHERE key = ANY($1::text[]);

-- name: GetPaymentMethodByID :one
SELECT *
FROM payment_methods
WHERE id = $1;