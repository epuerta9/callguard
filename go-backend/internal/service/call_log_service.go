package service

import (
	"context"
	"errors"

	"github.com/epuerta/callguard/go-backend/internal/model"
	"github.com/epuerta/callguard/go-backend/internal/repository"
)

// CallLogService handles call log business logic
type CallLogService struct {
	repo *repository.CallLogRepository
}

// NewCallLogService creates a new CallLogService
func NewCallLogService(repo *repository.CallLogRepository) *CallLogService {
	return &CallLogService{repo: repo}
}

// GetByID gets a call log by ID
func (s *CallLogService) GetByID(ctx context.Context, id string, userID string) (*model.CallLog, error) {
	callLog, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Check if the call log belongs to the user
	if callLog.UserID != userID {
		return nil, errors.New("call log not found")
	}

	return callLog, nil
}

// List lists call logs with pagination
func (s *CallLogService) List(ctx context.Context, limit, offset int, userID string) ([]*model.CallLog, error) {
	return s.repo.ListByUserID(ctx, userID, int32(limit), int32(offset))
}

// Create creates a new call log
func (s *CallLogService) Create(ctx context.Context, callLog *model.CreateCallLogRequest, userID string) (*model.CallLog, error) {
	return s.repo.Create(ctx, callLog, userID)
}

// Update updates a call log
func (s *CallLogService) Update(ctx context.Context, id string, callLog *model.UpdateCallLogRequest, userID string) (*model.CallLog, error) {
	// Check if the call log exists and belongs to the user
	_, err := s.GetByID(ctx, id, userID)
	if err != nil {
		return nil, err
	}

	// If we got here, the call log exists and belongs to the user
	return s.repo.Update(ctx, id, callLog)
}

// Delete deletes a call log
func (s *CallLogService) Delete(ctx context.Context, id string, userID string) error {
	// Check if the call log exists and belongs to the user
	_, err := s.GetByID(ctx, id, userID)
	if err != nil {
		return err
	}

	// If we got here, the call log exists and belongs to the user
	return s.repo.Delete(ctx, id)
}
