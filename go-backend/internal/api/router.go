package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/epuerta/callguard/go-backend/internal/config"
	"github.com/epuerta/callguard/go-backend/internal/model"
	"github.com/epuerta/callguard/go-backend/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewRouter creates a new API router
func NewRouter(cfg *config.Config, userService *service.UserService, callLogService *service.CallLogService) *echo.Echo {
	e := echo.New()

	// Create webhook service
	webhookService := service.NewWebhookService(callLogService)

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 60 * time.Second,
	}))

	// CORS configuration
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"}, // Replace with your frontend URL in production
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		ExposeHeaders:    []string{echo.HeaderContentLength},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not readily refused by browsers
	}))

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	})

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	})

	e.POST("/", func(c echo.Context) error {
		fmt.Println("Received webhook")

		// Parse the request body into a WebhookPayload
		var webhookPayload model.WebhookPayload
		if err := c.Bind(&webhookPayload); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Failed to parse webhook message",
			})
		}

		fmt.Printf("Webhook type: %s\n", webhookPayload.Message.Type)

		// Handle the webhook message based on its type
		if err := webhookService.HandleWebhook(webhookPayload.Message); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"status": "webhook processed successfully",
		})
	})

	// API routes
	api := e.Group("/api/v1")

	// Public routes
	api.POST("/users/register", userHandler(userService).register)
	api.POST("/users/login", userHandler(userService).login)

	// Protected routes
	protected := api.Group("")
	protected.Use(authMiddleware)

	// User routes
	protected.GET("/users/me", userHandler(userService).getCurrent)
	protected.PUT("/users/me", userHandler(userService).updateCurrent)

	// Call logs routes
	protected.GET("/call-logs", callLogHandler(callLogService).list)
	protected.POST("/call-logs", callLogHandler(callLogService).create)
	protected.GET("/call-logs/:id", callLogHandler(callLogService).get)
	protected.PUT("/call-logs/:id", callLogHandler(callLogService).update)
	protected.DELETE("/call-logs/:id", callLogHandler(callLogService).delete)

	return e
}
