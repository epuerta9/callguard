-- name: GetUserByID :one
SELECT * FROM users 
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = $1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name
LIMIT $1 OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (
  name, email, password_hash
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET name = $2,
    email = $3,
    password_hash = CASE WHEN $4 = '' THEN password_hash ELSE $4 END
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1; 