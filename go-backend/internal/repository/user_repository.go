package repository

import (
	"context"
	"time"

	"github.com/epuerta/callguard/go-backend/internal/db"
	"github.com/epuerta/callguard/go-backend/internal/model"
	"github.com/google/uuid"
)

// UserRepository provides access to the user data store
type UserRepository struct {
	db *db.Queries
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *db.Queries) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// GetByID retrieves a user by their ID
func (r *UserRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	// Convert id string to UUID
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	dbUser, err := r.db.GetUserByID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return convertDBUserToUser(dbUser), nil
}

// GetByEmail retrieves a user by their email
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	dbUser, err := r.db.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return convertDBUserToUser(dbUser), nil
}

// Create creates a new user
func (r *UserRepository) Create(ctx context.Context, req *model.RegisterRequest, passwordHash string) (*model.User, error) {
	// Placeholder implementation
	return &model.User{
		ID:           uuid.New().String(),
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: passwordHash,
		CreatedAt:    "2023-06-01T10:00:00Z",
		UpdatedAt:    "2023-06-01T10:00:00Z",
	}, nil
}

// Update updates an existing user
func (r *UserRepository) Update(ctx context.Context, id string, req *model.UpdateUserRequest, newPasswordHash string) (*model.User, error) {
	// Placeholder implementation
	passwordHash := "existing-hash"
	if newPasswordHash != "" {
		passwordHash = newPasswordHash
	}

	return &model.User{
		ID:           id,
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: passwordHash,
		CreatedAt:    "2023-06-01T10:00:00Z",
		UpdatedAt:    "2023-06-01T11:00:00Z",
	}, nil
}

// Delete deletes a user
func (r *UserRepository) Delete(ctx context.Context, id string) error {
	// Placeholder implementation
	return nil
}

// convertDBUserToUser converts a db.User to a model.User
func convertDBUserToUser(dbUser db.User) *model.User {
	return &model.User{
		ID:           dbUser.ID.String(),
		Name:         dbUser.Name,
		Email:        dbUser.Email,
		PasswordHash: dbUser.PasswordHash,
		CreatedAt:    dbUser.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:    dbUser.UpdatedAt.Time.Format(time.RFC3339),
	}
}
