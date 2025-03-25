package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/epuerta/callguard/go-backend/internal/api"
	"github.com/epuerta/callguard/go-backend/internal/config"
	"github.com/epuerta/callguard/go-backend/internal/db"
	"github.com/epuerta/callguard/go-backend/internal/repository"
	"github.com/epuerta/callguard/go-backend/internal/service"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	// Setup logger
	logger := setupLogger(cfg)
	slog.SetDefault(logger)

	// Connect to the database
	dbPool, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}
	defer dbPool.Close()

	// Run database migrations
	if err := db.RunMigrations(cfg); err != nil {
		slog.Error("Failed to run database migrations", "error", err)
		os.Exit(1)
	}

	// Initialize repositories
	queries := db.New(dbPool)
	userRepo := repository.NewUserRepository(queries)
	callLogRepo := repository.NewCallLogRepository(queries)

	// Initialize services
	userService := service.NewUserService(userRepo)
	callLogService := service.NewCallLogService(callLogRepo)

	// Initialize the API server
	e := api.NewRouter(cfg, userService, callLogService)

	// Start the server in a goroutine
	go func() {
		slog.Info("Starting server", "port", cfg.Port)
		if err := e.Start(":8000"); err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
		os.Exit(1)
	}

	slog.Info("Server exiting")
}

func setupLogger(cfg *config.Config) *slog.Logger {
	var logLevel slog.Level

	// Default to info level
	logLevel = slog.LevelInfo

	// Use development-friendly logging in development mode
	if cfg.Environment == "development" {
		return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
	}

	// Use JSON logging in production
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
}
