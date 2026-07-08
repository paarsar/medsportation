# Medsportation Website Makefile

# Project name
PROJECT_NAME=medsportation-wb

# Build directory (relative to this folder)
BUILD_DIR=dist/medsportation-wb/browser

.PHONY: all build deploy clean

# Default target
all: build

# Build the Angular project
build:
	@echo "Building the project..."
	npm install && npm run build

# Deploy to Cloudflare Pages
deploy: build
	@echo "Deploying to Cloudflare Pages..."
	npx wrangler pages deploy $(BUILD_DIR) --branch=production --project-name $(PROJECT_NAME)

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf dist
