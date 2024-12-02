# 📋 Go Todo List Web Application

## Overview

A lightweight, responsive web-based todo list application built with Go and Tailwind CSS. The application allows users to add, manage, and track their tasks with an intuitive interface.

## Features

- 🚀 Add new tasks with a single click
- ✅ Mark tasks as complete
- 📅 Automatic date tracking for each task
- 🎨 Responsive and modern user interface
- 🔒 Concurrent-safe task management

## Application Workflow

1. **Task Addition**
   - Enter a task in the input field
   - Click the "+" button to add the task
   - Task is automatically timestamped with the current date

2. **Task Management**
   - View all tasks in a clean, organized list
   - Checkbox next to each task for marking completion
   - Completed tasks are:
     * Visually crossed out
     * Moved to the bottom of the task list

## Prerequisites

- Docker
- Docker Compose (optional)
- Go 1.21+ (for local development)

## Running the Application

### Docker Deployment

#### Single-Stage Dockerfile

```dockerfile
FROM golang:1.21

WORKDIR /app
COPY . .

RUN go build -o todo-list

EXPOSE 8080
CMD ["./todo-list"]
```

#### Multi-Stage Dockerfile

```dockerfile
# Build Stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o todo-list

# Production Stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/todo-list .
COPY --from=builder /app/index.html .
EXPOSE 8080
CMD ["./todo-list"]
```

### Running with Docker

#### Build and Run

```bash
# Build the Docker image
docker build -t todo-list .

# Run the container
docker run -p 8080:8080 todo-list
```

### Docker Compose Setup

Create a `docker-compose.yml`:

```yaml
version: '3.8'

services:
  todo-app:
    build: .
    ports:
      - "8080:8080"
    volumes:
      - ./:/app
    restart: unless-stopped
```

Run with Docker Compose:
```bash
docker-compose up -d
```

## Local Development

### Prerequisites
- Go 1.21+
- Make (optional)

### Steps
1. Clone the repository
2. Install dependencies
```bash
go mod tidy
```
3. Run the application
```bash
go run main.go
```

## Application Interface

![Todo List Interface](interface-screenshot.png)

### Interface Components
- Minimalist design with a clean, modern look
- Input field for task entry
- "+" button for adding tasks
- Checkbox for task completion
- Responsive layout works on desktop and mobile

## Technical Details

- **Backend**: Go (Golang)
- **Frontend**: HTML5, Tailwind CSS
- **Concurrency**: `sync.RWMutex` for thread-safe operations
- **Sorting**: Custom algorithm to prioritize incomplete tasks

## Performance Considerations

- Lightweight and fast
- Minimal resource consumption
- Concurrent-safe task management
- Efficient task rendering

## Roadmap and Future Improvements

- [ ] Persistent storage (database integration)
- [ ] User authentication
- [ ] Task categorization
- [ ] Advanced filtering options

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License

This project is open-source. See the LICENSE file for details.

## Contact

For issues, feature requests, or contributions, please open an issue on GitHub.
