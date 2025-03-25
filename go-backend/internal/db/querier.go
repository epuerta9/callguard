// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateCallLog(ctx context.Context, arg CreateCallLogParams) (CallLog, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateVoiceAssistant(ctx context.Context, arg CreateVoiceAssistantParams) (VoiceAssistant, error)
	DeleteCallLog(ctx context.Context, id uuid.UUID) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	DeleteUserMetadataField(ctx context.Context, arg DeleteUserMetadataFieldParams) (User, error)
	DeleteVoiceAssistant(ctx context.Context, id uuid.UUID) error
	GetCallLogByID(ctx context.Context, id uuid.UUID) (CallLog, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (User, error)
	GetUserMetadata(ctx context.Context, id uuid.UUID) ([]byte, error)
	GetVoiceAssistantByID(ctx context.Context, id uuid.UUID) (VoiceAssistant, error)
	ListCallLogs(ctx context.Context, arg ListCallLogsParams) ([]CallLog, error)
	ListCallLogsByUserID(ctx context.Context, arg ListCallLogsByUserIDParams) ([]CallLog, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	ListVoiceAssistants(ctx context.Context, arg ListVoiceAssistantsParams) ([]VoiceAssistant, error)
	ListVoiceAssistantsByUserID(ctx context.Context, arg ListVoiceAssistantsByUserIDParams) ([]VoiceAssistant, error)
	SetUserMetadataField(ctx context.Context, arg SetUserMetadataFieldParams) (User, error)
	UpdateCallLog(ctx context.Context, arg UpdateCallLogParams) (CallLog, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateUserMetadata(ctx context.Context, arg UpdateUserMetadataParams) (User, error)
	UpdateVoiceAssistant(ctx context.Context, arg UpdateVoiceAssistantParams) (VoiceAssistant, error)
}

var _ Querier = (*Queries)(nil)
