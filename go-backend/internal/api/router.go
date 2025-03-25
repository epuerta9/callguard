package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	vapiclient "github.com/VapiAI/server-sdk-go/client"
	"github.com/VapiAI/server-sdk-go/option"
	"github.com/epuerta/callguard/go-backend/internal/config"
	"github.com/epuerta/callguard/go-backend/internal/db"
	"github.com/epuerta/callguard/go-backend/internal/model"
	"github.com/epuerta/callguard/go-backend/internal/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewRouter creates a new API router
func NewRouter(cfg *config.Config, userService *service.UserService, callLogService *service.CallLogService, queries *db.Queries) *echo.Echo {
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

	// Create user handler
	uh := userHandler(userService)

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

	// Auth routes
	e.POST("/auth/signup", uh.register)
	e.POST("/auth/login", uh.login)
	e.GET("/auth/me", uh.getCurrent)
	e.PUT("/auth/me", uh.updateCurrent)

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	})

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "healthy"})
	})

	e.POST("/", func(c echo.Context) error {
		fmt.Println("Received webhook")

		// Read the body once at the beginning
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			fmt.Println("Failed to read request body:", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Failed to read request body",
			})
		}
		// Restore the body for subsequent reads
		c.Request().Body = io.NopCloser(bytes.NewBuffer(body))

		// First try to bind to webhookPayload
		var webhookPayload model.WebhookPayload
		if err := json.Unmarshal(body, &webhookPayload); err != nil {
			fmt.Printf("Request body: %s\n", string(body))
			fmt.Println("Failed to parse webhook in router", err)
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Failed to parse webhook message",
			})
		}

		fmt.Printf("Received webhook message: %+v\n", webhookPayload.Message)
		if webhookPayload.Message.Type == "transfer-destination-request" {
			fmt.Println("Received transfer-destination-request webhook message")
			fmt.Printf("Transfer destination request payload: %s\n", string(body))

			var forwardPayload service.VAPIForwardPayload
			if err := json.Unmarshal(body, &forwardPayload); err != nil {
				fmt.Println("Failed to parse forward payload:", err)
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": "Failed to parse forward payload",
				})
			}
			fmt.Printf("Forward payload: %+v\n", forwardPayload)
			response, err := webhookService.HandleForwardWebhook(context.Background(), forwardPayload)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"error": err.Error(),
				})
			}
			return c.JSON(200, response)
		}
		// Handle the webhook message based on its type
		response, err := webhookService.HandleWebhook(webhookPayload.Message)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(200, response)
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
