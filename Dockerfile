# Use official golang image as base
FROM golang:1.24.3

# Set working directory
WORKDIR /app

# Copy entire project
COPY . .

# Download dependencies
RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main ./cmd/server/main.go

# Create non-root user for security
RUN adduser --system --no-create-home appuser

# Ensure the .env file is accessible
RUN chown appuser:root /app/.env
USER appuser

# Document the port the server listens on
EXPOSE 8080

# Set the startup command
CMD ["/app/main"]