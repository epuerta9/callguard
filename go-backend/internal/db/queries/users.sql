-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (
    name,
    email,
    password_hash
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET
    name = $2,
    email = $3,
    password_hash = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUserMetadata :one
UPDATE users
SET metadata = $2::text
WHERE id = $1
RETURNING *;

-- name: GetUserMetadata :one
SELECT metadata::text
FROM users
WHERE id = $1;

-- name: SetUserMetadataField :one
UPDATE users
SET metadata = $3::text
WHERE id = $1 AND $2::text IS NOT NULL -- $2 is the field name, kept for API compatibility
RETURNING *;

-- name: DeleteUserMetadataField :one
UPDATE users
SET metadata = $2::text
WHERE id = $1
RETURNING *; 