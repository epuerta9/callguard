package handlers

import (
	"net/http"

	"github.com/epuerta/callguard/go-backend/internal/model"
	"github.com/epuerta/callguard/go-backend/internal/repository"
	"github.com/labstack/echo/v4"
)

// TagHandler handles HTTP requests related to tags
type TagHandler struct {
	tagRepo *repository.TagRepository
}

// NewTagHandler creates a new TagHandler
func NewTagHandler(tagRepo *repository.TagRepository) *TagHandler {
	return &TagHandler{
		tagRepo: tagRepo,
	}
}

// RegisterRoutes registers the tag routes
func (h *TagHandler) RegisterRoutes(e *echo.Echo) {
	tags := e.Group("/api/tags")
	tags.GET("", h.ListTags)
	tags.POST("", h.CreateTag)
	tags.GET("/:id", h.GetTag)
	tags.PUT("/:id", h.UpdateTag)
	tags.DELETE("/:id", h.DeleteTag)
}

// ListTags handles GET /api/tags
func (h *TagHandler) ListTags(c echo.Context) error {
	tags, err := h.tagRepo.List(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve tags")
	}

	return c.JSON(http.StatusOK, tags)
}

// GetTag handles GET /api/tags/:id
func (h *TagHandler) GetTag(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing tag ID")
	}

	tag, err := h.tagRepo.GetByID(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Tag not found")
	}

	return c.JSON(http.StatusOK, tag)
}

// CreateTag handles POST /api/tags
func (h *TagHandler) CreateTag(c echo.Context) error {
	var req model.CreateTagRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}

	// TODO: Validate request

	tag, err := h.tagRepo.Create(c.Request().Context(), &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create tag")
	}

	return c.JSON(http.StatusCreated, tag)
}

// UpdateTag handles PUT /api/tags/:id
func (h *TagHandler) UpdateTag(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing tag ID")
	}

	var req model.UpdateTagRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request format")
	}

	// TODO: Validate request

	tag, err := h.tagRepo.Update(c.Request().Context(), id, &req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update tag")
	}

	return c.JSON(http.StatusOK, tag)
}

// DeleteTag handles DELETE /api/tags/:id
func (h *TagHandler) DeleteTag(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing tag ID")
	}

	err := h.tagRepo.Delete(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete tag")
	}

	return c.NoContent(http.StatusNoContent)
}
