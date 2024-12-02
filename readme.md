# Docker Learning Projects

## 🚀 Overview

This repository is a collection of small, focused Docker projects designed to help developers learn and explore containerization concepts. Each project serves as a practical example of different application types, deployment strategies, and Docker best practices.

## 📦 Project Collection

### 1. Age Calculator
- **Technology**: Go (Golang)
- **Description**: A simple web-based age calculator
- **Features**:
  * Calculates age based on birthdate
  * Responsive web interface
  * Multi-stage Docker build

### 2. Todo List Application
- **Technology**: Go (Golang)
- **Description**: A web-based task management application
- **Features**:
  * Add and manage tasks
  * Task completion tracking
  * Responsive design
  * Concurrent-safe task management

## 🛠 Prerequisites

- Docker
- Docker Compose
- Basic understanding of containerization concepts

## 💻 Getting Started

### Clone the Repository
```bash
git clone https://github.com/yourusername/docker-learning-projects.git
cd docker-learning-projects
```

### Building and Running Projects

#### Age Calculator
```bash
cd age-calculator
docker build -t age-calculator .
docker run -p 8080:8080 age-calculator
```

#### Todo List
```bash
cd todo-list
docker build -t todo-list .
docker run -p 8080:8080 todo-list
```

## 🌟 Key Learning Objectives

- Multi-stage Docker builds
- Optimizing container images
- Web application containerization
- Go language fundamentals
- Responsive web design
- Dockerizing different types of applications

## 📝 Best Practices Demonstrated

- Using Alpine-based images for minimal size
- Separating build and runtime environments
- Utilizing multi-stage builds
- Implementing concurrent-safe application logic
- Creating responsive web interfaces

## 🔧 Project Structure

Each project follows a consistent structure:
```
project-name/
│
├── main.go           # Main application logic
├── index.html        # Web interface
├── Dockerfile        # Docker build instructions
├── go.mod            # Go module definition
└── README.md         # Project-specific documentation
```

## 🚀 Deployment Considerations

- Each application is designed to be easily deployable
- Accessible on local network
- Minimal resource consumption
- Simple port mapping

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## 📄 License

This project is open-source. Please check the LICENSE file for details.

## 📚 Learning Resources

- [Docker Documentation](https://docs.docker.com/)
- [Go Programming Language](https://golang.org/doc/)
- [Containerization Fundamentals](https://www.docker.com/resources/what-container/)

## 🔍 Disclaimer

These projects are meant for learning purposes. They are simplified examples and should not be used directly in production without further refinement.
