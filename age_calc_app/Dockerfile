# Stage 1: Build Stage
FROM golang:1.21-alpine AS builder

# Set Working Directory
WORKDIR /app

# Copy Go mod and Sum Files
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o age-calculator -ldflags="-w -s" .

# Stage 2: Production stage
FROM scratch

# Set working directory
WORKDIR /root/

# Copy the built executable from the builder stage
COPY --from=builder /app/age-calculator .
COPY --from=builder /app/index.html .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./age-calculator"]
