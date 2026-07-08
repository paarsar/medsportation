# Medsportation Backend Makefile

APP_NAME=medsportation-be
REGION=ewr

.PHONY: all run build deploy logs status volume help

# Default target
all: help

## Local Development
run:
	@echo "Starting local development server..."
	go run cmd/main.go

build:
	@echo "Building local binary..."
	go build -o bin/main ./cmd/main.go

## Fly.io Deployment
deploy:
	@echo "Deploying to Fly.io..."
	fly deploy

logs:
	@echo "Streaming logs from Fly.io..."
	fly logs

status:
	@echo "Checking application status on Fly.io..."
	fly status

volume:
	@echo "Creating persistent volume for SQLite (one-time setup)..."
	fly volumes create sqlite_data --region $(REGION) --size 1

help:
	@echo "Medsportation Backend Management"
	@echo ""
	@echo "Local Commands:"
	@echo "  make run      - Run the Go server locally"
	@echo "  make build    - Build the Go binary locally"
	@echo ""
	@echo "Fly.io Commands:"
	@echo "  make deploy   - Deploy the application to Fly.io"
	@echo "  make logs     - View real-time application logs"
	@echo "  make status   - Check deployment status"
	@echo "  make volume   - Create the required SQLite volume (ewr region)"
	@echo "  make help     - Show this help message"
