-- name: GetVoiceAssistantByID :one
SELECT * FROM voice_assistants
WHERE id = $1;

-- name: ListVoiceAssistants :many
SELECT * FROM voice_assistants
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListVoiceAssistantsByUserID :many
SELECT * FROM voice_assistants
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: CreateVoiceAssistant :one
INSERT INTO voice_assistants (
    user_id,
    assistant_name,
    phone_number
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: UpdateVoiceAssistant :one
UPDATE voice_assistants
SET
    assistant_name = $2,
    phone_number = $3,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteVoiceAssistant :exec
DELETE FROM voice_assistants
WHERE id = $1; 