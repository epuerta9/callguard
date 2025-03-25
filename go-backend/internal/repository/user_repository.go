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
	// Create user in database
	dbUser, err := r.db.CreateUser(ctx, db.CreateUserParams{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: passwordHash,
	})
	if err != nil {
		return nil, err
	}

	return convertDBUserToUser(dbUser), nil
}

// Update updates an existing user
func (r *UserRepository) Update(ctx context.Context, id string, req *model.UpdateUserRequest, newPasswordHash string) (*model.User, error) {
	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	// Get existing user to preserve unchanged fields
	existingUser, err := r.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update only provided fields
	name := existingUser.Name
	if req.Name != "" {
		name = req.Name
	}

	passwordHash := existingUser.PasswordHash
	if newPasswordHash != "" {
		passwordHash = newPasswordHash
	}

	// Update user in database
	dbUser, err := r.db.UpdateUser(ctx, db.UpdateUserParams{
		ID:           userID,
		Name:         name,
		Email:        existingUser.Email,
		PasswordHash: passwordHash,
	})
	if err != nil {
		return nil, err
	}

	return convertDBUserToUser(dbUser), nil
}

// Delete deletes a user
func (r *UserRepository) Delete(ctx context.Context, id string) error {
	userID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return r.db.DeleteUser(ctx, userID)
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
