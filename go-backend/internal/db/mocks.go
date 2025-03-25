// This file provides mock database functionality for development
package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// Helper function to create a timestamp
func mockTimestamp(timeStr string) pgtype.Timestamptz {
	t, _ := time.Parse(time.RFC3339, timeStr)
	return pgtype.Timestamptz{Time: t, Valid: true}
}

// Helper function to create a text field
func mockText(s string) pgtype.Text {
	return pgtype.Text{String: s, Valid: true}
}

// Helper function to create a float8 field
func mockFloat8(f float64) pgtype.Float8 {
	return pgtype.Float8{Float64: f, Valid: true}
}

// GetUserByID retrieves a user by their ID (mock implementation)
func (q *Queries) MockGetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	return User{
		ID:           id,
		Name:         "John Doe",
		Email:        "john@example.com",
		PasswordHash: "hashed-password",
		CreatedAt:    mockTimestamp("2023-06-01T10:00:00Z"),
		UpdatedAt:    mockTimestamp("2023-06-01T10:00:00Z"),
	}, nil
}

// GetUserByEmail retrieves a user by their email (mock implementation)
func (q *Queries) MockGetUserByEmail(ctx context.Context, email string) (User, error) {
	return User{
		ID:           uuid.New(),
		Name:         "John Doe",
		Email:        email,
		PasswordHash: "hashed-password",
		CreatedAt:    mockTimestamp("2023-06-01T10:00:00Z"),
		UpdatedAt:    mockTimestamp("2023-06-01T10:00:00Z"),
	}, nil
}

// GetTagByID retrieves a tag by its ID (mock implementation)
func (q *Queries) MockGetTagByID(ctx context.Context, id uuid.UUID) (Tag, error) {
	return Tag{
		ID:        id,
		Name:      "Suspicious",
		CreatedAt: mockTimestamp("2023-06-01T10:00:00Z"),
	}, nil
}

// GetTagsForCallLog retrieves all tags for a call log (mock implementation)
func (q *Queries) MockGetTagsForCallLog(ctx context.Context, callLogID uuid.UUID) ([]Tag, error) {
	return []Tag{
		{
			ID:        uuid.New(),
			Name:      "Suspicious",
			CreatedAt: mockTimestamp("2023-06-01T10:00:00Z"),
		},
		{
			ID:        uuid.New(),
			Name:      "Fraudulent",
			CreatedAt: mockTimestamp("2023-06-01T10:05:00Z"),
		},
	}, nil
}

// GetCallLogByID retrieves a call log by its ID (mock implementation)
func (q *Queries) MockGetCallLogByID(ctx context.Context, id uuid.UUID) (CallLog, error) {
	return CallLog{
		ID:             id,
		UserID:         uuid.New(),
		CallerNumber:   "+1234567890",
		CallDuration:   120,
		RecordingUrl:   mockText("https://example.com/recordings/1.mp3"),
		Transcript:     mockText("Hello, this is a test call."),
		SentimentScore: mockFloat8(0.8),
		RiskScore:      mockFloat8(0.2),
		Flagged:        false,
		Notes:          mockText("This is a test call"),
		CreatedAt:      mockTimestamp("2023-06-01T10:00:00Z"),
		UpdatedAt:      mockTimestamp("2023-06-01T10:00:00Z"),
	}, nil
}
