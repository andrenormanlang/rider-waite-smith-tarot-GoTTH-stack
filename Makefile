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
APP_NAME=tarot-app

.PHONY: tailwind-watch
tailwind-watch:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

.PHONY: tailwind-build
tailwind-build:
	npx tailwindcss -i ./static/css/input.css -o ./static/css/style.css --minify

.PHONY: templ-generate
templ-generate:
	templ generate

.PHONY: dev
dev:
	go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && air

.PHONY: build
build:
	make tailwind-build && make templ-generate && go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go