# Stage 1: Build stage
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod file
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o todo-list -ldflags="-w -s" .

# Stage 2: Production stage
FROM alpine:latest

# Install CA certificates
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy the built executable and HTML template
COPY --from=builder /app/todo-list .
COPY --from=builder /app/index.html .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./todo-list"]
