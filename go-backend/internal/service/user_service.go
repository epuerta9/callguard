package service

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/epuerta/callguard/go-backend/internal/model"
	"github.com/epuerta/callguard/go-backend/internal/repository"
	"github.com/golang-jwt/jwt/v5"
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

	// Generate a JWT token
	token, err := generateToken(user.ID)
	if err != nil {
		return nil, err
	}

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

// generateToken generates a JWT token for the given user ID
func generateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
