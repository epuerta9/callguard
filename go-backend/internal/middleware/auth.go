package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/epuerta/callguard/go-backend/internal/model"
	"github.com/labstack/echo/v4"
)

// Key type for context values
type contextKey string

// ContextUserKey is the key for the user in the request context
const ContextUserKey contextKey = "user"

// AuthEcho authenticates the request using Echo
func AuthEcho(c echo.Context, next echo.HandlerFunc) error {
	// Get the Authorization header
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header required")
	}

	// Check if the header has the right format
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization format")
	}

	// Extract the token
	token := strings.TrimPrefix(authHeader, "Bearer ")
	fmt.Println("token", token)
	if token == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token required")
	}

	// Validate the token (in a real application, you would verify the JWT)
	// This is a simple placeholder
	user, err := validateToken(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("Invalid token: %v", err))
	}

	fmt.Println("user", user)

	// Add the user to the request context
	c.Set(string(ContextUserKey), user)
	return next(c)
}

// Auth authenticates the request for http.Handler middleware (used with chi)
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Check if the header has the right format
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}

		// Extract the token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			http.Error(w, "Token required", http.StatusUnauthorized)
			return
		}

		// Validate the token (in a real application, you would verify the JWT)
		// This is a simple placeholder
		user, err := validateToken(token)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid token: %v", err), http.StatusUnauthorized)
			return
		}

		// Add the user to the request context
		ctx := context.WithValue(r.Context(), ContextUserKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserFromContext gets the user from the request context
func GetUserFromContext(ctx context.Context) (*model.User, bool) {
	user, ok := ctx.Value(ContextUserKey).(*model.User)
	return user, ok
}

// GetUserFromEchoContext gets the user from the Echo context
func GetUserFromEchoContext(c echo.Context) (*model.User, bool) {
	user, ok := c.Get(string(ContextUserKey)).(*model.User)
	return user, ok
}

// validateToken validates the JWT token and returns the user
// In a real application, this would verify the JWT and fetch the user from the database
func validateToken(token string) (*model.User, error) {
	// This is a placeholder for JWT validation
	// In a real application, you would verify the JWT signature and decode the claims

	// For demonstration purposes, we'll return a dummy user
	// In a real application, you would query the database with the user ID from the token
	return &model.User{
		ID:    "2063b131-a0d0-4f02-b004-6293b7b8e6e8", // Placeholder UUID
		Name:  "Test User",
		Email: "test@example.com",
	}, nil
}
