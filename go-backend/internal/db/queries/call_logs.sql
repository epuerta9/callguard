-- name: GetCallLogByID :one
SELECT * FROM call_logs
WHERE id = $1;

-- name: ListCallLogs :many
SELECT * FROM call_logs
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListCallLogsByUserID :many
SELECT * FROM call_logs
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: CreateCallLog :one
INSERT INTO call_logs (
    user_id,
    voice_assistant_id,
    caller_number,
    call_duration,
    transcript,
    is_potentially_malicious
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: UpdateCallLog :one
UPDATE call_logs
SET
    caller_number = $2,
    call_duration = $3,
    transcript = $4,
    is_potentially_malicious = $5,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteCallLog :exec
DELETE FROM call_logs
WHERE id = $1; 