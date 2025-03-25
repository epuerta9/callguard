package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	vapiclient "github.com/VapiAI/server-sdk-go/client"
	"github.com/VapiAI/server-sdk-go/option"
	"github.com/epuerta/callguard/go-backend/internal/config"
	"github.com/epuerta/callguard/go-backend/internal/model"
	"github.com/epuerta/callguard/go-backend/internal/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewRouter creates a new API router
func NewRouter(cfg *config.Config, userService *service.UserService, callLogService *service.CallLogService) *echo.Echo {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	e := echo.New()

	// Create Vapi client
	vapiClient := vapiclient.NewClient(option.WithToken(os.Getenv("VAPI_API_KEY")))
	vapiService := service.NewVapiService(vapiClient)
	// Create webhook service
	webhookService := service.NewWebhookService(callLogService, vapiService)

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

		var webhookPayload model.WebhookPayload
		if err := c.Bind(&webhookPayload); err != nil {
			fmt.Println("Failed to parse webhook message", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Failed to parse webhook message",
			})
		}

		fmt.Printf("Received webhook message: %+v\n", webhookPayload.Message)

		// Handle the webhook message based on its type
		response, err := webhookService.HandleWebhook(webhookPayload.Message)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		if response != nil {

			return c.JSON(http.StatusOK, map[string]interface{}{
				"status":   "webhook processed successfully",
				"response": response,
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
