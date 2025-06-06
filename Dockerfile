
# Stage 1: Build stage
FROM golang:1.24.3-alpine AS builder

WORKDIR /app

# Copy only the files needed for downloading dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application with optimizations
RUN CGO_ENABLED=0 GOOS=linux \
    go build -ldflags="-w -s" \
    -v -o main ./cmd/server/main.go

# Stage 2: Final stage
FROM alpine:3.19

WORKDIR /app

# Add necessary packages for DNS resolution and certificates
RUN apk --no-cache add \
    ca-certificates \
    bind-tools \
    tzdata \
    && adduser --system --no-create-home appuser

# Copy only the necessary files from builder
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Set proper permissions
RUN chown -R appuser:root /app && \
    chmod 500 /app/main && \
    chmod 400 /app/.env

# Switch to non-root user
USER appuser

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8080/health || exit 1


CMD ["/app/main"]