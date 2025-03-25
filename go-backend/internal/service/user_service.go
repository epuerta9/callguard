package service

import (
	"context"
	"errors"
	"time"

	"github.com/epuerta/callguard/go-backend/internal/model"
	"github.com/epuerta/callguard/go-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

// UserService handles user-related business logic
type UserService struct {
	repo *repository.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// GetByID gets a user by ID
func (s *UserService) GetByID(ctx context.Context, id string) (*model.User, error) {
	return s.repo.GetByID(ctx, id)
}

// Register registers a new user
func (s *UserService) Register(ctx context.Context, params *model.RegisterRequest) (*model.User, error) {
	// Check if the email is already in use
	_, err := s.repo.GetByEmail(ctx, params.Email)
	if err == nil {
		return nil, errors.New("email already in use")
	}

	// Hash the password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create the user
	return s.repo.Create(ctx, params, string(passwordHash))
}

// Login logs in a user
func (s *UserService) Login(ctx context.Context, params *model.LoginRequest) (*model.LoginResponse, error) {
	// Get the user by email
	user, err := s.repo.GetByEmail(ctx, params.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Check the password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(params.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate a token (in a real application, you would generate a JWT)
	// This is a simple placeholder
	token := generateDummyToken(user.ID)

	return &model.LoginResponse{
		User:  user,
		Token: token,
	}, nil
}

// Update updates a user
func (s *UserService) Update(ctx context.Context, userID string, params *model.UpdateUserRequest) (*model.User, error) {
	// Get the current user
	currentUser, err := s.repo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// If changing password, verify the current password
	passwordHash := ""
	if params.NewPassword != "" {
		// Verify current password
		err = bcrypt.CompareHashAndPassword([]byte(currentUser.PasswordHash), []byte(params.CurrentPassword))
		if err != nil {
			return nil, errors.New("current password is incorrect")
		}

		// Hash the new password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		passwordHash = string(hashedPassword)
	}

	// Update the user
	return s.repo.Update(ctx, userID, params, passwordHash)
}

// Delete deletes a user
func (s *UserService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

// generateDummyToken generates a dummy token for demonstration purposes
// In a real application, you would generate a JWT with appropriate claims
func generateDummyToken(userID string) string {
	return "dummy_token_" + userID + "_" + time.Now().Format(time.RFC3339)
}
