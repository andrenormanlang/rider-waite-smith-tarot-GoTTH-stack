# # Go parameters
# GOCMD=go
# TEMPL=templ
# BUILD_DIR=./tmp
# TAROT_DIR=./cmd/tarot-app
# TAROT_ADMIN_DIR=./cmd/tarot-admin-app

# # Name of the binary
# BINARY_NAME=tarot.exe
# ADMIN_BINARY_NAME=tarot-admin

# all: build test

# build:
# 	$(TEMPL) generate
# 	GIN_MODE=release $(GOCMD) build -ldflags "-s" -v -o $(BUILD_DIR)/$(BINARY_NAME) $(TAROT_DIR)
# 	GIN_MODE=release $(GOCMD) build -ldflags "-s" -v -o $(BUILD_DIR)/$(ADMIN_BINARY_NAME) $(TAROT_ADMIN_DIR)

# test:
# 	$(GOCMD) test -v ./...

# clean:
# 	$(GOCMD) clean
# 	rm -rf $(BUILD_DIR)

# install-tools:
# 	go install github.com/a-h/templ/cmd/templ@v0.2.543

# .PHONY: all build test clean


# Go parameters
APP_NAME=cmd/tarot-app
ADMIN_NAME=cmd/tarot-admin-app
BUILD_DIR=./tmp

# Load environment variables from .env file
export $(shell sed 's/=.*//' .env)

.PHONY: clean
clean:
	@rm -f ./tmp/app-main.exe ./tmp/admin-main.exe
	@echo "Cleaned old binaries."

.PHONY: tailwind-watch
tailwind-watch:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/style.css --minify

.PHONY: templ-generate
templ-generate:
	templ generate

# Build the frontend (tarot-app) and run it with Air (live reloading)
.PHONY: dev-app
dev-app:
	@echo "Running frontend app..."
	@go build -o $(BUILD_DIR)/app-main.exe ./cmd/tarot-app/main.go
	@air -c app.air.toml

# Build the backend (tarot-admin-app) and run it with Air (live reloading)
.PHONY: dev-admin
dev-admin:
	@echo "Running admin app..."
	@go build -o $(BUILD_DIR)/admin-main.exe ./cmd/tarot-admin-app/main.go
	@air -c admin.air.toml

# Run both frontend and backend apps in parallel
.PHONY: dev-all
dev-all:
	$(MAKE) clean
	$(MAKE) dev-app & 
	$(MAKE) dev-admin

# Build both frontend and backend apps for production
.PHONY: build
build:
	$(MAKE) tailwind-build && $(MAKE) templ-generate
	@go build -o $(BUILD_DIR)/app-main.exe ./cmd/tarot-app/main.go
	@go build -o $(BUILD_DIR)/admin-main.exe ./cmd/tarot-admin-app/main.go

# Clean up binaries and temporary files
.PHONY: clean
clean:
	@rm -rf $(BUILD_DIR)/*
	@echo "Cleaned up build files."
