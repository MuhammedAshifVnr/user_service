# Build Stage
FROM golang:1.22.2-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application binary
RUN go build -o user-svc ./cmd

# Runtime Stage
FROM alpine:3.18

# Set the working directory for the runtime container
WORKDIR /app

# Copy the built binary and environment file from the build stage
COPY --from=builder /app/user-svc .
COPY --from=builder /app/.env ./

# Expose the application port
EXPOSE 5001

# Start the application
CMD ["./user-svc"]