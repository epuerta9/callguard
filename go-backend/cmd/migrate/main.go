package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib" // PostgreSQL driver
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Get database URL from environment
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	// Set up command
	if len(os.Args) < 2 {
		log.Fatal("Command required: up, down, create")
	}
	command := os.Args[1]

	// Set the database driver
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Failed to set dialect: %v", err)
	}

	// Establish a database connection for migrations
	db, err := goose.OpenDBWithDriver("pgx", dbURL)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer db.Close()

	// Get the absolute path to the migrations directory
	migrationsDir := "internal/db/migrations"
	if !filepath.IsAbs(migrationsDir) {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Failed to get working directory: %v", err)
		}
		migrationsDir = filepath.Join(wd, migrationsDir)
	}

	// Execute the command
	switch command {
	case "up":
		if err := goose.Up(db, migrationsDir); err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}
		fmt.Println("Migrations applied successfully!")

	case "down":
		if err := goose.Down(db, migrationsDir); err != nil {
			log.Fatalf("Failed to roll back migrations: %v", err)
		}
		fmt.Println("Migrations rolled back successfully!")

	case "create":
		if len(os.Args) < 3 {
			log.Fatal("Migration name required")
		}
		name := strings.ToLower(strings.Join(os.Args[2:], "_"))
		if err := goose.Create(db, migrationsDir, name, "sql"); err != nil {
			log.Fatalf("Failed to create migration: %v", err)
		}
		fmt.Printf("Migration file created in %s\n", migrationsDir)

	default:
		log.Fatalf("Unknown command: %s", command)
	}
}
