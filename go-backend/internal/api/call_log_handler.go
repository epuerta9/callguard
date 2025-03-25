package api

import (
	"net/http"
	"strconv"

	"github.com/epuerta/callguard/go-backend/internal/middleware"
	"github.com/epuerta/callguard/go-backend/internal/model"
	"github.com/epuerta/callguard/go-backend/internal/service"
	"github.com/labstack/echo/v4"
)

type callLogHandlerImpl struct {
	service *service.CallLogService
}

func callLogHandler(service *service.CallLogService) *callLogHandlerImpl {
	return &callLogHandlerImpl{service: service}
}

func (h *callLogHandlerImpl) list(c echo.Context) error {
	user, ok := middleware.GetUserFromEchoContext(c)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "User not authenticated")
	}

	// Get pagination params
	limit := 10
	offset := 0

	limitParam := c.QueryParam("limit")
	if limitParam != "" {
		parsedLimit, err := strconv.Atoi(limitParam)
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	offsetParam := c.QueryParam("offset")
	if offsetParam != "" {
		parsedOffset, err := strconv.Atoi(offsetParam)
		if err == nil && parsedOffset >= 0 {
			offset = parsedOffset
		}
	}

	callLogs, err := h.service.List(c.Request().Context(), limit, offset, user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, callLogs)
}

func (h *callLogHandlerImpl) get(c echo.Context) error {
	user, ok := middleware.GetUserFromEchoContext(c)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "User not authenticated")
	}

	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Call log ID is required")
	}

	callLog, err := h.service.GetByID(c.Request().Context(), id, user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Call log not found")
	}

	return c.JSON(http.StatusOK, callLog)
}

func (h *callLogHandlerImpl) create(c echo.Context) error {
	user, ok := middleware.GetUserFromEchoContext(c)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "User not authenticated")
	}

	var req model.CreateCallLogRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	callLog, err := h.service.Create(c.Request().Context(), &req, user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, callLog)
}

func (h *callLogHandlerImpl) update(c echo.Context) error {
	user, ok := middleware.GetUserFromEchoContext(c)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "User not authenticated")
	}

	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Call log ID is required")
	}

	var req model.UpdateCallLogRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	callLog, err := h.service.Update(c.Request().Context(), id, &req, user.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, callLog)
}

func (h *callLogHandlerImpl) delete(c echo.Context) error {
	user, ok := middleware.GetUserFromEchoContext(c)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "User not authenticated")
	}

	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Call log ID is required")
	}

	if err := h.service.Delete(c.Request().Context(), id, user.ID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
