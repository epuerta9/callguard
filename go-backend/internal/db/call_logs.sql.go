// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: call_logs.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const addTagToCallLog = `-- name: AddTagToCallLog :exec
INSERT INTO call_logs_tags (call_log_id, tag_id)
VALUES ($1, $2)
ON CONFLICT DO NOTHING
`

type AddTagToCallLogParams struct {
	CallLogID uuid.UUID `json:"call_log_id"`
	TagID     uuid.UUID `json:"tag_id"`
}

func (q *Queries) AddTagToCallLog(ctx context.Context, arg AddTagToCallLogParams) error {
	_, err := q.db.Exec(addTagToCallLog, arg.CallLogID, arg.TagID)
	return err
}

const createCallLog = `-- name: CreateCallLog :one
INSERT INTO call_logs (
  user_id, caller_number, call_duration, recording_url, transcript, 
  sentiment_score, risk_score, flagged, notes
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING id, user_id, caller_number, call_duration, recording_url, transcript, sentiment_score, risk_score, flagged, notes, created_at, updated_at
`

type CreateCallLogParams struct {
	UserID         uuid.UUID     `json:"user_id"`
	CallerNumber   string        `json:"caller_number"`
	CallDuration   int32         `json:"call_duration"`
	RecordingUrl   pgtype.Text   `json:"recording_url"`
	Transcript     pgtype.Text   `json:"transcript"`
	SentimentScore pgtype.Float8 `json:"sentiment_score"`
	RiskScore      pgtype.Float8 `json:"risk_score"`
	Flagged        bool          `json:"flagged"`
	Notes          pgtype.Text   `json:"notes"`
}

func (q *Queries) CreateCallLog(ctx context.Context, arg CreateCallLogParams) (CallLog, error) {
	row := q.db.QueryRow(createCallLog,
		arg.UserID,
		arg.CallerNumber,
		arg.CallDuration,
		arg.RecordingUrl,
		arg.Transcript,
		arg.SentimentScore,
		arg.RiskScore,
		arg.Flagged,
		arg.Notes,
	)
	var i CallLog
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CallerNumber,
		&i.CallDuration,
		&i.RecordingUrl,
		&i.Transcript,
		&i.SentimentScore,
		&i.RiskScore,
		&i.Flagged,
		&i.Notes,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteCallLog = `-- name: DeleteCallLog :exec
DELETE FROM call_logs
WHERE id = $1
`

func (q *Queries) DeleteCallLog(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(deleteCallLog, id)
	return err
}

const getCallLogByID = `-- name: GetCallLogByID :one
SELECT id, user_id, caller_number, call_duration, recording_url, transcript, sentiment_score, risk_score, flagged, notes, created_at, updated_at FROM call_logs 
WHERE id = $1
`

func (q *Queries) GetCallLogByID(ctx context.Context, id uuid.UUID) (CallLog, error) {
	row := q.db.QueryRow(getCallLogByID, id)
	var i CallLog
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CallerNumber,
		&i.CallDuration,
		&i.RecordingUrl,
		&i.Transcript,
		&i.SentimentScore,
		&i.RiskScore,
		&i.Flagged,
		&i.Notes,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTagsForCallLog = `-- name: GetTagsForCallLog :many
SELECT t.id, t.name, t.created_at FROM tags t
JOIN call_logs_tags clt ON t.id = clt.tag_id
WHERE clt.call_log_id = $1
ORDER BY t.name
`

func (q *Queries) GetTagsForCallLog(ctx context.Context, callLogID uuid.UUID) ([]Tag, error) {
	rows, err := q.db.Query(getTagsForCallLog, callLogID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tag{}
	for rows.Next() {
		var i Tag
		if err := rows.Scan(&i.ID, &i.Name, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCallLogs = `-- name: ListCallLogs :many
SELECT id, user_id, caller_number, call_duration, recording_url, transcript, sentiment_score, risk_score, flagged, notes, created_at, updated_at FROM call_logs
ORDER BY created_at DESC
LIMIT $1 OFFSET $2
`

type ListCallLogsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCallLogs(ctx context.Context, arg ListCallLogsParams) ([]CallLog, error) {
	rows, err := q.db.Query(listCallLogs, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CallLog{}
	for rows.Next() {
		var i CallLog
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CallerNumber,
			&i.CallDuration,
			&i.RecordingUrl,
			&i.Transcript,
			&i.SentimentScore,
			&i.RiskScore,
			&i.Flagged,
			&i.Notes,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listCallLogsByUserID = `-- name: ListCallLogsByUserID :many
SELECT id, user_id, caller_number, call_duration, recording_url, transcript, sentiment_score, risk_score, flagged, notes, created_at, updated_at FROM call_logs
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type ListCallLogsByUserIDParams struct {
	UserID uuid.UUID `json:"user_id"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

func (q *Queries) ListCallLogsByUserID(ctx context.Context, arg ListCallLogsByUserIDParams) ([]CallLog, error) {
	rows, err := q.db.Query(listCallLogsByUserID, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CallLog{}
	for rows.Next() {
		var i CallLog
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CallerNumber,
			&i.CallDuration,
			&i.RecordingUrl,
			&i.Transcript,
			&i.SentimentScore,
			&i.RiskScore,
			&i.Flagged,
			&i.Notes,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeTagFromCallLog = `-- name: RemoveTagFromCallLog :exec
DELETE FROM call_logs_tags
WHERE call_log_id = $1 AND tag_id = $2
`

type RemoveTagFromCallLogParams struct {
	CallLogID uuid.UUID `json:"call_log_id"`
	TagID     uuid.UUID `json:"tag_id"`
}

func (q *Queries) RemoveTagFromCallLog(ctx context.Context, arg RemoveTagFromCallLogParams) error {
	_, err := q.db.Exec(removeTagFromCallLog, arg.CallLogID, arg.TagID)
	return err
}

const updateCallLog = `-- name: UpdateCallLog :one
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
RETURNING id, user_id, caller_number, call_duration, recording_url, transcript, sentiment_score, risk_score, flagged, notes, created_at, updated_at
`

type UpdateCallLogParams struct {
	ID             uuid.UUID     `json:"id"`
	CallerNumber   string        `json:"caller_number"`
	CallDuration   int32         `json:"call_duration"`
	RecordingUrl   pgtype.Text   `json:"recording_url"`
	Transcript     pgtype.Text   `json:"transcript"`
	SentimentScore pgtype.Float8 `json:"sentiment_score"`
	RiskScore      pgtype.Float8 `json:"risk_score"`
	Flagged        bool          `json:"flagged"`
	Notes          pgtype.Text   `json:"notes"`
}

func (q *Queries) UpdateCallLog(ctx context.Context, arg UpdateCallLogParams) (CallLog, error) {
	row := q.db.QueryRow(updateCallLog,
		arg.ID,
		arg.CallerNumber,
		arg.CallDuration,
		arg.RecordingUrl,
		arg.Transcript,
		arg.SentimentScore,
		arg.RiskScore,
		arg.Flagged,
		arg.Notes,
	)
	var i CallLog
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CallerNumber,
		&i.CallDuration,
		&i.RecordingUrl,
		&i.Transcript,
		&i.SentimentScore,
		&i.RiskScore,
		&i.Flagged,
		&i.Notes,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
