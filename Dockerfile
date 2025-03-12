# Step 1: Build the Go application
FROM golang:latest AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go application
RUN GOOS=linux GOARCH=amd64 go build -o main ./cmd/server

# Step 2: Create the final image
FROM alpine:latest

# Install dependencies like bash and ca-certificates
RUN apk --no-cache add ca-certificates

# Copy the compiled binary from the builder image
COPY --from=builder /app/main /app/

# Expose the application port
EXPOSE 8080

# Set the entry point to run the Go binary
CMD ["/app/main"]
