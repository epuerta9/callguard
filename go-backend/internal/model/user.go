package model

import (
	"encoding/json"
	"time"
)

// User represents a user in the system
type User struct {
	ID           string          `json:"id"`
	Name         string          `json:"name"`
	Email        string          `json:"email"`
	PasswordHash string          `json:"-"` // Never expose in JSON
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	Metadata     json.RawMessage `json:"metadata,omitempty"`
}

// LoginRequest represents a login request
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// RegisterRequest represents a registration request
type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// UpdateUserRequest represents a request to update a user
type UpdateUserRequest struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	CurrentPassword string `json:"current_password" validate:"omitempty,min=8"`
	NewPassword     string `json:"new_password" validate:"omitempty,min=8"`
}

// LoginResponse represents a login response
type LoginResponse struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}
