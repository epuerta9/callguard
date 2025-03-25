package repository

import (
	"context"
	"encoding/json"

	"github.com/epuerta/callguard/go-backend/internal/db"
	"github.com/epuerta/callguard/go-backend/internal/model"
	"github.com/google/uuid"
)

// UserRepository handles database operations for users
type UserRepository struct {
	db *db.Queries
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *db.Queries) *UserRepository {
	return &UserRepository{db: db}
}

// GetByID gets a user by ID
func (r *UserRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	user, err := r.db.GetUserByID(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return convertDBUserToUser(user), nil
}

// GetByEmail gets a user by email
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := r.db.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return convertDBUserToUser(user), nil
}

// Create creates a new user
func (r *UserRepository) Create(ctx context.Context, params *model.RegisterRequest, passwordHash string) (*model.User, error) {
	user, err := r.db.CreateUser(ctx, db.CreateUserParams{
		Name:         params.Name,
		Email:        params.Email,
		PasswordHash: passwordHash,
	})
	if err != nil {
		return nil, err
	}
	return convertDBUserToUser(user), nil
}

// Update updates a user
func (r *UserRepository) Update(ctx context.Context, id string, params *model.UpdateUserRequest, passwordHash string) (*model.User, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	user, err := r.db.UpdateUser(ctx, db.UpdateUserParams{
		ID:           uuid,
		Name:         params.Name,
		Email:        params.Email,
		PasswordHash: passwordHash,
	})
	if err != nil {
		return nil, err
	}
	return convertDBUserToUser(user), nil
}

// Delete deletes a user
func (r *UserRepository) Delete(ctx context.Context, id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return r.db.DeleteUser(ctx, uuid)
}

// UpdateMetadata updates the user's metadata
func (r *UserRepository) UpdateMetadata(ctx context.Context, userID string, metadata json.RawMessage) (*model.User, error) {
	uuid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	user, err := r.db.UpdateUserMetadata(ctx, db.UpdateUserMetadataParams{
		ID:       uuid,
		Metadata: metadata,
	})
	if err != nil {
		return nil, err
	}
	return convertDBUserToUser(user), nil
}

// GetMetadata gets the user's metadata
func (r *UserRepository) GetMetadata(ctx context.Context, userID string) (json.RawMessage, error) {
	uuid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	metadata, err := r.db.GetUserMetadata(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}

// SetMetadataField sets a specific field in the user's metadata
func (r *UserRepository) SetMetadataField(ctx context.Context, userID string, field string, value json.RawMessage) (*model.User, error) {
	uuid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	user, err := r.db.SetUserMetadataField(ctx, db.SetUserMetadataFieldParams{
		ID:      uuid,
		Column2: field,
		Column3: value,
	})
	if err != nil {
		return nil, err
	}
	return convertDBUserToUser(user), nil
}

// DeleteMetadataField removes a field from the user's metadata
func (r *UserRepository) DeleteMetadataField(ctx context.Context, userID string, field string) (*model.User, error) {
	uuid, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	user, err := r.db.DeleteUserMetadataField(ctx, db.DeleteUserMetadataFieldParams{
		ID:       uuid,
		Metadata: []byte(field),
	})
	if err != nil {
		return nil, err
	}
	return convertDBUserToUser(user), nil
}

// convertDBUserToUser converts a db.User to a model.User
func convertDBUserToUser(dbUser db.User) *model.User {
	return &model.User{
		ID:           dbUser.ID.String(),
		Name:         dbUser.Name,
		Email:        dbUser.Email,
		PasswordHash: dbUser.PasswordHash,
		CreatedAt:    dbUser.CreatedAt.Time,
		UpdatedAt:    dbUser.UpdatedAt.Time,
		Metadata:     dbUser.Metadata,
	}
}
