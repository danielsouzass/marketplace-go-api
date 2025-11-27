-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens ("expires_in", "user_id")
VALUES ($1, $2)
    RETURNING *;

-- name: GetRefreshTokenByUserID :one
SELECT *
FROM refresh_tokens
WHERE user_id = $1;

-- name: DeleteRefreshToken :exec
DELETE FROM refresh_tokens
WHERE id = $1;