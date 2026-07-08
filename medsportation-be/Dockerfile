# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary (disable CGO for glebarez/sqlite)
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Run stage
FROM alpine:latest

WORKDIR /app

# Install CA certificates for SMTP (SSL/TLS) and sqlite for maintenance
RUN apk --no-cache add ca-certificates sqlite

# Copy the binary
COPY --from=builder /app/main .

# Create a directory for the database volume
RUN mkdir -p /data

# Set environment variables for production
ENV GIN_MODE=release
ENV DATABASE_PATH=/data/medsportation.db

EXPOSE 8080

CMD ["./main"]
