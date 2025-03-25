package db

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/epuerta/callguard/go-backend/internal/config"
	"github.com/pressly/goose/v3"
)

// RunMigrations runs database migrations
func RunMigrations(cfg *config.Config) error {
	// Set the database driver
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set dialect: %w", err)
	}

	// Establish a database connection for migrations
	db, err := goose.OpenDBWithDriver("postgres", cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}
	defer db.Close()

	// Get the absolute path to the migrations directory
	migrationsDir := "internal/db/migrations"
	if !filepath.IsAbs(migrationsDir) {
		wd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get working directory: %w", err)
		}
		migrationsDir = filepath.Join(wd, migrationsDir)
	}

	// Run migrations
	if err := goose.Up(db, migrationsDir); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}
