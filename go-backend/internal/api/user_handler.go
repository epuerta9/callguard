package api

import (
	"net/http"

	"github.com/epuerta/callguard/go-backend/internal/middleware"
	"github.com/epuerta/callguard/go-backend/internal/model"
	"github.com/epuerta/callguard/go-backend/internal/service"
	"github.com/labstack/echo/v4"
)

type userHandlerImpl struct {
	service *service.UserService
}

func userHandler(service *service.UserService) *userHandlerImpl {
	return &userHandlerImpl{service: service}
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

	resp, err := h.service.Login(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
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
