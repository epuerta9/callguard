# CallGuard Go Backend

This is the Go backend for the CallGuard application. It provides a RESTful API for user management and call logging.

## Technologies Used

- Go (1.22+)
- PostgreSQL
- Chi router for HTTP routing
- pgx/v5 for PostgreSQL connectivity
- sqlc for type-safe SQL queries
- Goose for database migrations

## Getting Started

### Prerequisites

- Go 1.22 or higher
- PostgreSQL
- Make (for using the Makefile commands)

### Setting Up the Development Environment

1. Clone the repository
2. Install the required tools:

```bash
make install-tools
```

3. Create a `.env` file in the root directory with the following contents:

```env
# Environment (development or production)
ENVIRONMENT=development
PORT=8080
LOG_LEVEL=debug

# Database configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=callguard
DB_SSL_MODE=disable
```

4. Initialize the database:

```bash
make db-init
```

5. Run the migrations:

```bash
make migrate-up
```

6. Generate the sqlc code:

```bash
make generate-sqlc
```

### Running the Application

For development with hot reloading:

```bash
make dev
```

For a standard run:

```bash
make run
```

### Building the Application

```bash
make build
```

This will create a binary in the `./bin` directory.

## API Endpoints

### Authentication

- `POST /api/v1/users/register` - Register a new user
- `POST /api/v1/users/login` - Login and get a token

### Users

- `GET /api/v1/users/me` - Get the current user profile
- `PUT /api/v1/users/me` - Update the current user profile

### Call Logs

- `GET /api/v1/call-logs` - List call logs for the current user
- `POST /api/v1/call-logs` - Create a new call log
- `GET /api/v1/call-logs/{id}` - Get a specific call log
- `PUT /api/v1/call-logs/{id}` - Update a call log
- `DELETE /api/v1/call-logs/{id}` - Delete a call log

## Development

### Creating a Migration

```bash
make migrate-create
```

You'll be prompted for a migration name.

### Database Migrations

Up:

```bash
make migrate-up
```

Down:

```bash
make migrate-down
```

### Testing

```bash
make test
```

## Docker

### Building Docker Image

```bash
make docker-build
```

### Running with Docker

```bash
make docker-run
```

## Project Structure

```
go-backend/
├── cmd/                  # Command-line applications
│   └── api/              # Main API application
├── internal/             # Private application code
│   ├── api/              # API handlers and routing
│   ├── config/           # Configuration handling
│   ├── db/               # Database access and migrations
│   │   ├── migrations/   # Database migration files
│   │   └── queries/      # SQLC query files
│   ├── middleware/       # HTTP middleware
│   ├── model/            # Data models
│   ├── repository/       # Database repositories
│   └── service/          # Business logic
└── Makefile              # Project tasks
``` 