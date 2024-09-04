# Use an official Go image as a base image
FROM golang:1.22.3-alpine AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files for dependency resolution
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code of your application
COPY . .

# Build Tarot app (frontend)
RUN go build -o ./tmp/app-main ./cmd/tarot-app/main.go

# Build Tarot Admin app (backend)
RUN go build -o ./tmp/admin-main ./cmd/tarot-admin-app/main.go

# Use a minimal base image for the final stage to reduce size
FROM alpine:latest

# Set the Current Working Directory
WORKDIR /app

# Install tini (a tiny init system to manage multiple processes)
RUN apk add --no-cache tini

# Copy the built binaries from the build stage
COPY --from=build /app/tmp/app-main /app/tmp/admin-main /app/tmp/

# Copy the .env file
COPY .env /app/.env

# Expose the necessary ports
EXPOSE 8080 8081

# Copy over static assets (if any)
COPY static/ /app/static
COPY images/ /app/images

# Use Tini as the entrypoint to manage both processes
ENTRYPOINT ["/sbin/tini", "--"]

# Command to run both the Tarot app (frontend) and Tarot Admin app (backend)
CMD ["/bin/sh", "-c", "/app/tmp/app-main & /app/tmp/admin-main"]
