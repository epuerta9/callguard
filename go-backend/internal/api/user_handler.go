package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/epuerta/callguard/go-backend/internal/db"
	"github.com/epuerta/callguard/go-backend/internal/middleware"
	"github.com/epuerta/callguard/go-backend/internal/model"
	"github.com/epuerta/callguard/go-backend/internal/service"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

type userHandlerImpl struct {
	service   *service.UserService
	queries   *db.Queries
	jwtSecret []byte
}

func userHandler(service *service.UserService, queries *db.Queries) *userHandlerImpl {
	return &userHandlerImpl{service: service, queries: queries}
}

func (h *userHandlerImpl) register(c echo.Context) error {
	var req model.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	user, err := h.service.Register(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *userHandlerImpl) login(c echo.Context) error {
	var req model.LoginRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	fmt.Println("Login", req)
	resp, err := h.service.Login(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *userHandlerImpl) Signup(c echo.Context) error {
	fmt.Println("Signup")
	var req model.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// if err := c.Validate(&req); err != nil {
	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// }

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to hash password")
	}

	// Create user
	user, err := h.queries.CreateUser(c.Request().Context(), db.CreateUserParams{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create user")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(h.jwtSecret)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
	}

	return c.JSON(http.StatusCreated, AuthResponse{
		AccessToken: tokenString,
	})
}

func (h *userHandlerImpl) getCurrent(c echo.Context) error {
	user, ok := middleware.GetUserFromEchoContext(c)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "User not authenticated")
	}

	return c.JSON(http.StatusOK, user)
}

func (h *userHandlerImpl) updateCurrent(c echo.Context) error {
	user, ok := middleware.GetUserFromEchoContext(c)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "User not authenticated")
	}

	var req model.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	updatedUser, err := h.service.Update(c.Request().Context(), user.ID, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, updatedUser)
}

// updateCurrentMetadata updates the current user's metadata
func (h *userHandlerImpl) updateCurrentMetadata(c echo.Context) error {
	user, ok := middleware.GetUserFromEchoContext(c)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "User not authenticated")
	}

	var metadata json.RawMessage
	if err := c.Bind(&metadata); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid metadata format")
	}
	fmt.Println("metadata", string(metadata))
	fmt.Println("user.ID", user.ID)
	updatedUser, err := h.service.UpdateMetadata(c.Request().Context(), user.ID, metadata)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, updatedUser)
}

// getCurrentMetadata gets the current user's metadata
func (h *userHandlerImpl) getCurrentMetadata(c echo.Context) error {
	user, ok := middleware.GetUserFromEchoContext(c)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "User not authenticated")
	}

	metadata, err := h.service.GetMetadata(c.Request().Context(), user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]json.RawMessage{
		"metadata": metadata,
	})
}

// setCurrentMetadataField sets a specific field in the current user's metadata
func (h *userHandlerImpl) setCurrentMetadataField(c echo.Context) error {
	user, ok := middleware.GetUserFromEchoContext(c)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "User not authenticated")
	}

	field := c.Param("field")
	fmt.Println("field", field)
	if field == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Field name is required")
	}

	var value json.RawMessage
	if err := c.Bind(&value); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid value format")
	}

	updatedUser, err := h.service.SetMetadataField(c.Request().Context(), user.ID, field, value)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, updatedUser)
}

// deleteCurrentMetadataField deletes a specific field from the current user's metadata
func (h *userHandlerImpl) deleteCurrentMetadataField(c echo.Context) error {
	user, ok := middleware.GetUserFromEchoContext(c)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "User not authenticated")
	}

	field := c.Param("field")
	if field == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Field name is required")
	}

	updatedUser, err := h.service.DeleteMetadataField(c.Request().Context(), user.ID, field)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, updatedUser)
}
