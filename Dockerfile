# --- Stage 1: Build the Go binary ---
    FROM golang:1.24-alpine AS builder

    # Set environment variables
    ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
    
    # Create working directory
    WORKDIR /app
    
    # Copy go mod and sum files
    COPY go.mod go.sum ./
    
    # Download dependencies
    RUN go mod download
    
    # Copy the rest of the source code
    COPY . .
    
    # Build the application
    RUN go build -o app ./cmd/event-processor
    
    # --- Stage 2: Create a minimal image to run the app ---
    FROM alpine:latest
    
    # Create a non-root user (optional but recommended)
    RUN adduser -D myuser
    
    # Copy the built binary from the builder stage
    COPY --from=builder /app/app /app/app
    
    # Use non-root user
    USER myuser
    
    # Set the working directory
    WORKDIR /app

    EXPOSE 3000
    
    # Command to run the executable
    CMD ["./app"]
    