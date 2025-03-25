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
  user_id, caller_number, call_duration, recording_url, transcript, 
  sentiment_score, risk_score, flagged, notes
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: UpdateCallLog :one
UPDATE call_logs
SET caller_number = $2,
    call_duration = $3,
    recording_url = $4,
    transcript = $5,
    sentiment_score = $6,
    risk_score = $7,
    flagged = $8,
    notes = $9
WHERE id = $1
RETURNING *;

-- name: DeleteCallLog :exec
DELETE FROM call_logs
WHERE id = $1;

-- name: AddTagToCallLog :exec
INSERT INTO call_logs_tags (call_log_id, tag_id)
VALUES ($1, $2)
ON CONFLICT DO NOTHING;

-- name: RemoveTagFromCallLog :exec
DELETE FROM call_logs_tags
WHERE call_log_id = $1 AND tag_id = $2;

-- name: GetTagsForCallLog :many
SELECT t.* FROM tags t
JOIN call_logs_tags clt ON t.id = clt.tag_id
WHERE clt.call_log_id = $1
ORDER BY t.name; 