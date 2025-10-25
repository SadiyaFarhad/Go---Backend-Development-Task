# ---------- STEP 1: Build the Go binary ----------
FROM golang:1.25.3-alpine AS builder

# Set working directory inside build stage
WORKDIR /app

# Copy Go modules and download dependencies first (for cache efficiency)
COPY go.mod go.sum ./
RUN go mod download

# Copy all project files
COPY . .

# Build the Go binary from cmd/server/main.go
RUN go build -o server ./cmd/server

# ---------- STEP 2: Create the minimal runtime image ----------
FROM alpine:latest

# Set working directory for the runtime container
WORKDIR /app

# Copy the compiled binary from builder stage
COPY --from=builder /app/server .

# Copy any required config or migration files if needed
COPY config ./config
COPY db ./db

# Expose application port
EXPOSE 3000

# Command to run your app
CMD ["./server"]
