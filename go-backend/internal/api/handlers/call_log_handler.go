package handlers

import (
	"net/http"
	"strconv"

	"github.com/epuerta/callguard/go-backend/internal/model"
	"github.com/epuerta/callguard/go-backend/internal/repository"
	"github.com/labstack/echo/v4"
)

// CallLogHandler handles HTTP requests related to call logs
type CallLogHandler struct {
	callLogRepo *repository.CallLogRepository
	tagRepo     *repository.TagRepository
}

// NewCallLogHandler creates a new CallLogHandler
func NewCallLogHandler(callLogRepo *repository.CallLogRepository, tagRepo *repository.TagRepository) *CallLogHandler {
	return &CallLogHandler{
		callLogRepo: callLogRepo,
		tagRepo:     tagRepo,
	}
}

// RegisterRoutes registers the call log routes
func (h *CallLogHandler) RegisterRoutes(e *echo.Echo) {
	callLogs := e.Group("/api/call-logs")
	callLogs.GET("", h.ListCallLogs)
	callLogs.POST("", h.CreateCallLog)
	callLogs.GET("/:id", h.GetCallLog)
	callLogs.PUT("/:id", h.UpdateCallLog)
	callLogs.DELETE("/:id", h.DeleteCallLog)
	callLogs.GET("/:id/tags", h.GetCallLogTags)
}

// ListCallLogs handles GET /api/call-logs
func (h *CallLogHandler) ListCallLogs(c echo.Context) error {
	userID := c.QueryParam("user_id")

	// Parse pagination parameters
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit < 1 || limit > 100 {
		limit = 20 // Default limit
	}

	var callLogs []*model.CallLog
	var err error

	if userID != "" {
		callLogs, err = h.callLogRepo.ListByUserID(c.Request().Context(), userID, int32(page), int32(limit))
	} else {
		callLogs, err = h.callLogRepo.List(c.Request().Context(), int32(page), int32(limit))
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve call logs")
	}

	return c.JSON(http.StatusOK, callLogs)
}

// GetCallLog handles GET /api/call-logs/:id
func (h *CallLogHandler) GetCallLog(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing call log ID")
	}

	callLog, err := h.callLogRepo.GetByID(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Call log not found")
	}

	return c.JSON(http.StatusOK, callLog)
}

// CreateCallLog handles POST /api/call-logs
func (h *CallLogHandler) CreateCallLog(c echo.Context) error {
	var req model.CreateCallLogRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}

	// TODO: Validate request

	// Get the user ID from the authenticated user or request
	userID := c.Get("userID").(string) // Assuming middleware sets this
	if userID == "" {
		userID = c.QueryParam("user_id") // Fallback to query param
		if userID == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "User ID is required")
		}
	}

	callLog, err := h.callLogRepo.Create(c.Request().Context(), &req, userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create call log")
	}

	return c.JSON(http.StatusCreated, callLog)
}

// UpdateCallLog handles PUT /api/call-logs/:id
func (h *CallLogHandler) UpdateCallLog(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing call log ID")
	}

	var req model.UpdateCallLogRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}

	// TODO: Validate request

	callLog, err := h.callLogRepo.Update(c.Request().Context(), id, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update call log")
	}

	return c.JSON(http.StatusOK, callLog)
}

// DeleteCallLog handles DELETE /api/call-logs/:id
func (h *CallLogHandler) DeleteCallLog(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing call log ID")
	}

	err := h.callLogRepo.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete call log")
	}

	return c.NoContent(http.StatusNoContent)
}

// GetCallLogTags handles GET /api/call-logs/:id/tags
func (h *CallLogHandler) GetCallLogTags(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing call log ID")
	}

	tags, err := h.tagRepo.GetTagsForCallLog(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve tags")
	}

	return c.JSON(http.StatusOK, tags)
}
