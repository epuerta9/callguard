.PHONY: build clean test generate-sqlc migrate-up migrate-down migrate-create dev run

# Build the application
build:
	go build -o ./bin/api ./cmd/api

# Clean build artifacts
clean:
	rm -rf ./bin

# Run tests
test:
	go test -v ./...

# Generate SQLC code
generate-sqlc:
	sqlc generate

# Run database migrations up
migrate-up:
	go run ./cmd/migrate/main.go up

# Run database migrations down
migrate-down:
	go run ./cmd/migrate/main.go down

# Create a new migration file
migrate-create:
	@read -p "Enter migration name: " name; \
	go run ./cmd/migrate/main.go create $$name

# Run development server with hot reload
dev:
	go run github.com/cosmtrek/air -c .air.toml

# Run the application
run:
	go run ./cmd/api/main.go

# Generate swagger documentation
swagger:
	swag init -g cmd/api/main.go -o api/docs

# Install development tools
install-tools:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/cosmtrek/air@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest

# Initialize a new database
db-init:
	@read -p "Enter database name (default: callguard): " dbname; \
	if [ -z "$$dbname" ]; then dbname="callguard"; fi; \
	createdb $$dbname || echo "Database already exists or error occurred"; \
	echo "Database $$dbname created or already exists"

# Docker commands
docker-build:
	docker build -t callguard-api .

docker-run:
	docker run -p 8080:8080 --env-file .env callguard-api

# Help command
help:
	@echo "Make commands:"
	@echo "  build          - Build the application"
	@echo "  clean          - Clean build artifacts"
	@echo "  test           - Run tests"
	@echo "  generate-sqlc  - Generate SQLC code"
	@echo "  migrate-up     - Run database migrations up"
	@echo "  migrate-down   - Run database migrations down"
	@echo "  migrate-create - Create a new migration"
	@echo "  dev            - Run development server with hot reload"
	@echo "  run            - Run the application"
	@echo "  swagger        - Generate swagger documentation"
	@echo "  install-tools  - Install development tools"
	@echo "  db-init        - Initialize a new database"
	@echo "  docker-build   - Build Docker image"
	@echo "  docker-run     - Run Docker container" 