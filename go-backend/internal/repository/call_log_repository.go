package repository

import (
	"context"
	"time"

	"github.com/epuerta/callguard/go-backend/internal/db"
	"github.com/epuerta/callguard/go-backend/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// CallLogRepository provides access to the call log data store
type CallLogRepository struct {
	db *db.Queries
}

// NewCallLogRepository creates a new CallLogRepository
func NewCallLogRepository(db *db.Queries) *CallLogRepository {
	return &CallLogRepository{
		db: db,
	}
}

// GetByID retrieves a call log by its ID
func (r *CallLogRepository) GetByID(ctx context.Context, id string) (*model.CallLog, error) {
	// Convert id string to UUID
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	dbLog, err := r.db.GetCallLogByID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return convertDBCallLogToCallLog(dbLog), nil
}

// List retrieves all call logs with pagination
func (r *CallLogRepository) List(ctx context.Context, page, limit int32) ([]*model.CallLog, error) {
	// Placeholder implementation
	callLogs := []*model.CallLog{
		{
			ID:             uuid.New().String(),
			UserID:         uuid.New().String(),
			CallerNumber:   "+1234567890",
			CallDuration:   120,
			RecordingURL:   "https://example.com/recordings/1.mp3",
			Transcript:     "Hello, this is a test call.",
			SentimentScore: 0.8,
			RiskScore:      0.2,
			Flagged:        false,
			Notes:          "This is a test call",
			Tags:           nil,
			CreatedAt:      "2023-06-01T10:00:00Z",
			UpdatedAt:      "2023-06-01T10:00:00Z",
		},
		{
			ID:             uuid.New().String(),
			UserID:         uuid.New().String(),
			CallerNumber:   "+0987654321",
			CallDuration:   180,
			RecordingURL:   "https://example.com/recordings/2.mp3",
			Transcript:     "Hello, this is another test call.",
			SentimentScore: 0.5,
			RiskScore:      0.6,
			Flagged:        true,
			Notes:          "This call seems suspicious",
			Tags:           nil,
			CreatedAt:      "2023-06-01T11:00:00Z",
			UpdatedAt:      "2023-06-01T11:00:00Z",
		},
	}
	return callLogs, nil
}

// ListByUserID retrieves all call logs for a user with pagination
func (r *CallLogRepository) ListByUserID(ctx context.Context, userID string, page, limit int32) ([]*model.CallLog, error) {
	// Placeholder implementation
	callLogs := []*model.CallLog{
		{
			ID:             uuid.New().String(),
			UserID:         userID,
			CallerNumber:   "+1234567890",
			CallDuration:   120,
			RecordingURL:   "https://example.com/recordings/1.mp3",
			Transcript:     "Hello, this is a test call.",
			SentimentScore: 0.8,
			RiskScore:      0.2,
			Flagged:        false,
			Notes:          "This is a test call",
			Tags:           nil,
			CreatedAt:      "2023-06-01T10:00:00Z",
			UpdatedAt:      "2023-06-01T10:00:00Z",
		},
		{
			ID:             uuid.New().String(),
			UserID:         userID,
			CallerNumber:   "+0987654321",
			CallDuration:   180,
			RecordingURL:   "https://example.com/recordings/2.mp3",
			Transcript:     "Hello, this is another test call.",
			SentimentScore: 0.5,
			RiskScore:      0.6,
			Flagged:        true,
			Notes:          "This call seems suspicious",
			Tags:           nil,
			CreatedAt:      "2023-06-01T11:00:00Z",
			UpdatedAt:      "2023-06-01T11:00:00Z",
		},
	}
	return callLogs, nil
}

// Create creates a new call log
func (r *CallLogRepository) Create(ctx context.Context, req *model.CreateCallLogRequest, userID string) (*model.CallLog, error) {
	// Convert userID to UUID
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	// Create the call log
	dbLog, err := r.db.CreateCallLog(ctx, db.CreateCallLogParams{
		UserID:       userUUID,
		CallerNumber: req.CallerNumber,
		CallDuration: int32(req.CallDuration),
	})
	if err != nil {
		return nil, err
	}

	// Get the created call log with tags
	return r.GetByID(ctx, dbLog.ID.String())
}

// Update updates an existing call log
func (r *CallLogRepository) Update(ctx context.Context, id string, req *model.UpdateCallLogRequest) (*model.CallLog, error) {
	// Convert id to UUID
	callLogID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	// Update the call log
	_, err = r.db.UpdateCallLog(ctx, db.UpdateCallLogParams{
		ID:           callLogID,
		CallerNumber: req.CallerNumber,
		CallDuration: int32(req.CallDuration),
		Transcript:   pgtype.Text{String: req.Transcript, Valid: req.Transcript != ""},
	})
	if err != nil {
		return nil, err
	}

	// Get the updated call log with tags
	return r.GetByID(ctx, id)
}

// Delete deletes a call log
func (r *CallLogRepository) Delete(ctx context.Context, id string) error {
	// Placeholder implementation
	return nil
}

// convertDBCallLogToCallLog converts a db.CallLog to a model.CallLog
func convertDBCallLogToCallLog(dbLog db.CallLog) *model.CallLog {
	return &model.CallLog{
		ID:           dbLog.ID.String(),
		UserID:       dbLog.UserID.String(),
		CallerNumber: dbLog.CallerNumber,
		CallDuration: int(dbLog.CallDuration),
		Transcript:   dbLog.Transcript.String,
		Tags:         nil,
		CreatedAt:    dbLog.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:    dbLog.UpdatedAt.Time.Format(time.RFC3339),
	}
}

// nullableFloat64 provides a helper for handling nullable float64 values
func nullableFloat64(v pgtype.Float8) float64 {
	if !v.Valid {
		return 0
	}
	return v.Float64
}

// pointerFloat64 provides a helper for handling *float64 values
func pointerFloat64(v *float64) float64 {
	if v == nil {
		return 0
	}
	return *v
}
