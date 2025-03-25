package handlers

import (
	"net/http"

	"github.com/epuerta/callguard/go-backend/internal/model"
	"github.com/epuerta/callguard/go-backend/internal/repository"
	"github.com/labstack/echo/v4"
)

// UserHandler handles HTTP requests related to users
type UserHandler struct {
	userRepo *repository.UserRepository
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

// RegisterRoutes registers the user routes
func (h *UserHandler) RegisterRoutes(e *echo.Echo) {
	auth := e.Group("/api/auth")
	auth.POST("/register", h.Register)
	auth.POST("/login", h.Login)

	users := e.Group("/api/users")
	users.GET("/me", h.GetCurrentUser) // Protected route
	users.PUT("/me", h.UpdateUser)     // Protected route
}

// Register handles POST /api/auth/register
func (h *UserHandler) Register(c echo.Context) error {
	var req model.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}

	// TODO: Validate request

	// Generate password hash
	passwordHash := "hashed-password" // Replace with actual hashing

	user, err := h.userRepo.Create(c.Request().Context(), &req, passwordHash)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
	}

	// Generate JWT token
	// TODO: Implement JWT token generation

	response := &model.LoginResponse{
		User:  user,
		Token: "jwt-token-placeholder", // Replace with actual JWT token
	}

	return c.JSON(http.StatusCreated, response)
}

// Login handles POST /api/auth/login
func (h *UserHandler) Login(c echo.Context) error {
	var req model.LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}

	// TODO: Validate request

	user, err := h.userRepo.GetByEmail(c.Request().Context(), req.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
	}

	// TODO: Verify password

	// Generate JWT token
	// TODO: Implement JWT token generation

	response := &model.LoginResponse{
		User:  user,
		Token: "jwt-token-placeholder", // Replace with actual JWT token
	}

	return c.JSON(http.StatusOK, response)
}

// GetCurrentUser handles GET /api/users/me
func (h *UserHandler) GetCurrentUser(c echo.Context) error {
	// Get user ID from JWT token
	// TODO: Implement JWT token extraction
	userID := c.Get("userID").(string) // Assuming middleware sets this

	user, err := h.userRepo.GetByID(c.Request().Context(), userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

// UpdateUser handles PUT /api/users/me
func (h *UserHandler) UpdateUser(c echo.Context) error {
	// Get user ID from JWT token
	// TODO: Implement JWT token extraction
	userID := c.Get("userID").(string) // Assuming middleware sets this

	var req model.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}

	// TODO: Validate request

	// Generate new password hash if needed
	var newPasswordHash string
	if req.NewPassword != "" {
		newPasswordHash = "hashed-new-password" // Replace with actual hashing
	}

	user, err := h.userRepo.Update(c.Request().Context(), userID, &req, newPasswordHash)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update user")
	}

	return c.JSON(http.StatusOK, user)
}
