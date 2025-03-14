# Step 1: Build the Go application
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go mod and sum files to cache dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go application for ARM-based architecture (arm64)
RUN GOOS=linux GOARCH=arm64 go build -o main ./cmd/server

# Step 2: Create the final image
FROM alpine:latest

# Install necessary dependencies for Go binary compatibility
RUN apk --no-cache add ca-certificates libc6-compat bash

# Copy the compiled binary from the builder image
COPY --from=builder /app/main /app/

# Ensure the binary has executable permissions
RUN chmod +x /app/main

# Expose the application port
EXPOSE 8080

# Set the entry point to run the Go binary
CMD ["/app/main"]
