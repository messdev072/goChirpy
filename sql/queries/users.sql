-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: UpdateUserByID :one
UPDATE users SET email = $2,
updated_at = NOW(),hashed_password=$3
WHERE id = $1
RETURNING *;

-- name: UpdateUserChirpyByID :one
UPDATE users SET is_chirpy_red = true
WHERE id = $1
RETURNING *;