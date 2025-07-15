## Goal Tracker

## Overview

This is a Goal Tracker developed in Go, designed to monitor objectives and their respective status. The application allows users to create, update, and track their goals with progress indicators.

The front-end is located in another repository called goal-tracker-front.

## Technologies Used

- Go 1.24.3
- PostgreSQL 16
- Docker & Docker Compose
- GORM (ORM for Go)
- Gorilla Mux (HTTP router)
- Swagger (API documentation)
- JWT for authentication (planned)

## Features

- User registration and authentication
- Create, read, update, and delete goals
- Track goal progress
- Health check endpoint for monitoring
- API rate limiting
- CORS and security headers
- Swagger documentation

## Requirements

- Docker and Docker Compose
- Go 1.24.3 or higher (for local development)
- Make (optional, for simplified commands)

## Environment Setup

1. Clone the repository:
```
git clone https://github.com/hvmello/goal-tracker-backend.git
cd goal-tracker-backend
```

2. Configure environment variables:
```
cp .env.example .env
```
Edit .env file with your settings. Available configuration options:

| Variable | Description | Default |
|----------|-------------|---------|
| SERVER_PORT | Port for the HTTP server | 8080 |
| SERVER_BASE_URL | Base URL for the server | "" |
| DB_HOST | Database host | localhost |
| DB_PORT | Database port | 5432 |
| DB_USER | Database username | postgres |
| DB_PASSWORD | Database password | postgres |
| DB_NAME | Database name | goaltracker |
| DB_SSL_MODE | SSL mode for database connection | disable |
| RATE_LIMIT_ENABLED | Enable rate limiting | true |
| RATE_LIMIT_REQUESTS_PER_MINUTE | Requests allowed per minute | 60 |
| RATE_LIMIT_BURST_SIZE | Burst size for rate limiting | 20 |

3. Start containers:
```
docker-compose up --build
```

## API Endpoints

### Health Check
`GET /health`

Checks the status of the application and its services.

**Success Response:**
```json
{ 
  "success": true, 
  "data": { 
    "status": "ok", 
    "timestamp": "2025-06-02T10:00:00Z", 
    "services": { 
      "database": "ok", 
      "memory": "24.5MB", 
      "goroutines": "8" 
    }, 
    "db_stats": { 
      "open_connections": 2, 
      "in_use": 1, 
      "idle": 1 
    }, 
    "version": "1.0.0" 
  } 
}
```

### Goals API

#### Get All Goals
`GET /api/goals`

Retrieves all goals for the authenticated user.

#### Get Goal by ID
`GET /api/goals/{id}`

Retrieves a specific goal by its ID.

#### Create Goal
`POST /api/goals`

Creates a new goal.

**Request Body:**
```json
{
  "title": "Learn Go Programming",
  "description": "Complete a Go course and build a project",
  "dueDate": "2023-12-31T23:59:59Z",
  "progress": 0,
  "userId": 1
}
```

#### Update Goal
`PUT /api/goals/{id}`

Updates an existing goal.

**Request Body:**
```json
{
  "title": "Learn Go Programming",
  "description": "Complete a Go course and build a project",
  "dueDate": "2023-12-31T23:59:59Z",
  "progress": 50,
  "userId": 1
}
```

#### Delete Goal
`DELETE /api/goals/{id}`

Deletes a goal by its ID.

### User API

#### Register User
`POST /api/users/register`

Registers a new user.

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "securepassword"
}
```

#### Login
`POST /api/users/login`

Authenticates a user.

**Request Body:**
```json
{
  "email": "john@example.com",
  "password": "securepassword"
}
```

## Project Structure

The project follows a clean architecture approach with the following structure:

```plaintext
├── cmd/                 # Main applications for this project
│   └── main.go          # Application entry point
├── docs/                # Swagger documentation
├── internal/            # Private application and library code
│   ├── config/          # Configuration handling
│   ├── goals/           # Goals domain
│   ├── health/          # Health check functionality
│   ├── middleware/      # HTTP middleware
│   ├── router/          # HTTP router setup
│   └── user/            # User domain
├── pkg/                 # Library code that can be used by external applications
│   └── response/        # HTTP response utilities
├── tests/               # Additional external test apps and test data
├── Dockerfile           # Docker configuration
├── docker-compose.yml   # Docker Compose configuration
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
└── README.md            # Project documentation
```


## Development Workflow

### Branch Naming Pattern with Semantic Versioning

The project follows a structured branch naming convention based on semantic versioning:

| Type            | Prefix + version                | Example                               |
| --------------- | ------------------------------ | ------------------------------------|
| Feature         | `feature/1.0.0/`               | `feature/1.0.0/create-goals-api`     |
| Bug fix         | `fix/1.0.1/`                   | `fix/1.0.1/database-connection`      |
| Infrastructure  | `chore/1.1.0/`                 | `chore/1.1.0/docker-setup`            |
| Refactoring     | `refactor/1.2.0/`              | `refactor/1.2.0/handler-cleanup`     |
| Urgent hotfix   | `hotfix/1.0.2/`                | `hotfix/1.0.2/fix-prod-crash`        |

### Tips for Using This Pattern

- Version `1.0.0` follows [Semantic Versioning (SemVer)](https://semver.org/), meaning:  
  - **MAJOR** (1): incompatible API changes  
  - **MINOR** (0): backward-compatible feature additions  
  - **PATCH** (0): backward-compatible bug fixes

- Update the version in the branch name as development progresses:  
  - Increment **MINOR** for new features without breaking changes  
  - Increment **PATCH** for bug fixes or small improvements  
  - Increment **MAJOR** for breaking changes

- Including versions in branch names helps with automation, tagging, and release management.

## Useful Commands

```bash
# Build and start containers
docker-compose up --build

# Start containers in background
docker-compose up -d

# Stop containers
docker-compose down

# Run tests
go test ./...

# Run tests with coverage
go test ./... -cover

# Check code
go vet ./...

# Generate Swagger documentation
swag init -g cmd/main.go -o docs

# Access Swagger UI
# Open http://localhost:8080/swagger/index.html in your browser
```

## Contributing

1. Fork the repository
2. Create a new branch following the branch naming pattern
3. Make your changes
4. Write or update tests as needed
5. Ensure all tests pass
6. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
