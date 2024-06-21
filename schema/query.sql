-- name: GetUserByEmail :one
SELECT * 
FROM users
WHERE email = ? 
LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    email, password
) VALUES (
    ?, ?
)
RETURNING *;
