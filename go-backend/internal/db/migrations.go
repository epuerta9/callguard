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
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}
	// If we're running from cmd/api, we need to go up two levels
	if filepath.Base(wd) == "api" && filepath.Base(filepath.Dir(wd)) == "cmd" {
		wd = filepath.Dir(filepath.Dir(wd))
	}
	migrationsDir := filepath.Join(wd, "internal", "db", "migrations")

	// Run migrations
	if err := goose.Up(db, migrationsDir); err != nil {
		// If the goose_db_version table already exists, we can ignore that error
		if err.Error() != "ERROR: relation \"goose_db_version\" already exists (SQLSTATE 42P07)" {
			return fmt.Errorf("failed to run migrations: %w", err)
		}
	}

	return nil
}
