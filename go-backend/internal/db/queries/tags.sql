-- name: GetTagByID :one
SELECT * FROM tags 
WHERE id = $1;

-- name: GetTagByName :one
SELECT * FROM tags 
WHERE name = $1;

-- name: ListTags :many
SELECT * FROM tags
ORDER BY name
LIMIT $1 OFFSET $2;

-- name: CreateTag :one
INSERT INTO tags (name)
VALUES ($1)
RETURNING *;

-- name: UpdateTag :one
UPDATE tags
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTag :exec
DELETE FROM tags
WHERE id = $1; 