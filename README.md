# REST-APP

A REST API application built with Go (Golang), featuring microservices architecture and Docker support.

## Features

- **Event Management**:
  - Create events
  - Update events
  - Delete events
  - Fetch all events
  
- **User Management**:
  - Create users
  - Authenticate users (JWT-based login)
  
- **Middleware**:
  - Authentication and authorization using JWT

## Technologies Used

- **Programming Language**: [Go (Golang)](https://golang.org/)
- **Database**: PostgreSQL (via `docker-compose`)
- **Containerization**: Docker
- **Frameworks/Libraries**:
  - [Gin](https://github.com/gin-gonic/gin) for building HTTP web services
  - [JWT](https://github.com/golang-jwt/jwt) for authentication

## Getting Started

### Prerequisites

Ensure you have the following installed:
- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### Installation

1. Clone the repository:
```bash
 git clone https://github.com/your-username/rest-app.git
 cd rest-app
 ```
2.Build and run the application using Docker Compose:
  ```bash
   docker-compose up --build
   ```
3.The API will be available at:
  ```bash
   http://localhost:8080
   ```

## Project Structure

```bash
.
├─ cmd/
│  └─ main.go           # Entry point of the application
├─ db/
│  ├─ db.go             # Database connection and initialization
│  └─ ...
├─ middleware/
│  └─ auth.go           # Authentication middleware (JWT handling)
├─ models/
│  ├─ user.go           # User model definition
│  ├─ event.go          # Event model definition
│  └─ ...
├─ routers/
│  ├─ user.go           # User-related HTTP endpoints
│  ├─ event.go          # Event-related HTTP endpoints
│  └─ ...
├─ utils/
│  └─ jwt.go            # JWT utility functions
├─ api-test/
│  ├─ create-event.http # Example test requests for creating events
│  ├─ get-events.http   # Example test requests for retrieving events
│  └─ ...               # Additional .http files for testing endpoints
├─ Dockerfile
├─ docker-compose.yml
├─ go.mod / go.sum
└─ README.md
```
