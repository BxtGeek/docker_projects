# 🧮 Go Age Calculator

## Overview

The Go Age Calculator is a simple, responsive web application that allows users to calculate their exact age by inputting their birthdate. Built with Go and Tailwind CSS, the application provides a clean and intuitive interface for calculating age with precision.

## Features

- 📅 Simple date input interface
- 🔢 Precise age calculation (years, months, days)
- 🌐 Responsive web design
- 🐳 Docker support for easy deployment
- 💻 Local network accessibility

## Screenshot

![Age Calculator Interface](https://placeholder.co/400x300?text=Age+Calculator+Screenshot)

*Note: Replace the placeholder with an actual screenshot of your application*

## Prerequisites

- Go 1.21+
- Docker (optional)
- Web browser

## Running the Application

### Local Development

1. Clone the repository:
```bash
git clone https://github.com/yourusername/age-calculator.git
cd age-calculator
```

2. Run the application:
```bash
go mod tidy
go run main.go
```

3. Open your browser and navigate to:
- `http://localhost:8080`

### Docker Deployment

#### Single-Stage Dockerfile

```dockerfile
FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go build -o age-calculator

EXPOSE 8080

CMD ["./age-calculator"]
```

#### Multi-Stage Dockerfile

```dockerfile
# Stage 1: Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o age-calculator

# Stage 2: Production stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/age-calculator .
COPY --from=builder /app/index.html .
EXPOSE 8080
CMD ["./age-calculator"]
```

### Docker Run Options

#### Direct Docker Run
```bash
# Build the image
docker build -t age-calculator .

# Run the container
docker run -p 8080:8080 age-calculator
```

#### Docker Compose

Create a `docker-compose.yml`:
```yaml
version: '3.8'

services:
  age-calculator:
    build: .
    ports:
      - "8080:8080"
    networks:
      - local-network

networks:
  local-network:
    driver: bridge
```

Run with Docker Compose:
```bash
docker-compose up -d
```

## Network Access

The application is configured to be accessible:
- Locally: `http://localhost:8080`
- On local network: `http://[YOUR_LOCAL_IP]:8080`

## How to Use

1. Open the Age Calculator in your web browser
2. Click on the date input field
3. Select your birthdate
4. Click "Calculate Age"
5. View your precise age in years, months, and days

## Technology Stack

- 🟢 Backend: Go (Golang)
- 🎨 Frontend: HTML, Tailwind CSS
- 🐳 Containerization: Docker

## Performance and Security

- Lightweight Go binary
- Minimal Alpine Linux base image
- Statically linked executable
- Secure, concurrent-safe design

## Troubleshooting

- Ensure port 8080 is not in use
- Check Docker installation
- Verify Go version compatibility

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License

[Choose an appropriate license, e.g., MIT]

## Contact

[Your Name]
[Your Email or GitHub Profile]
