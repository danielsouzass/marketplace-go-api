-- name: GetPaymentMethodsByKeys :many
SELECT id, key, name
FROM payment_methods
WHERE key = ANY($1::text[]);

-- name: ListPaymentMethods :many
SELECT *
FROM payment_methods
ORDER BY name;