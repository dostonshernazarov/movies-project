# Movies APP

A RESTful API for managing movie records built with Go, Gin, GORM, and UberFx.

## Features

- CRUD operations for movies
- JWT authentication and authorization
- Transaction handling
- Input validation
- Error handling
- Swagger documentation
- Docker support

## Technologies

- Go
- Gin (HTTP routing)
- GORM (ORM with PostgreSQL)
- UberFx (Dependency Injection)
- JWT (Authentication)
- Docker & Docker Compose
- Swagger (API Documentation)

## Project Structure

```
movies-project/
├── config/             # Configuration
├── controllers/        # HTTP request handlers
├── core/               # Application core
├── docs/               # Swagger documentation
├── middleware/         # HTTP middleware
├── models/             # Database models
├── repositories/       # Data access layer
├── services/           # Business logic
├── .env                # Environment variables (not in git)
├── .env.example        # Example environment variables
├── Dockerfile          # Docker configuration
├── docker-compose.yml  # Docker Compose configuration
├── go.mod              # Go modules
├── go.sum              # Go modules checksums
├── main.go             # Application entry point
└── README.md           # Project documentation
```

## Getting Started

### Prerequisites

- Go 1.20 or higher
- PostgreSQL
- Docker & Docker Compose (optional)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/dostonshernazarov/movies-project.git
   cd movies-project
   ```

2. Copy the example environment file:
   ```bash
   cp .env.example .env
   ```

3. Modify the `.env` file with your configuration.

### Running with Go

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

### Running with Docker

1. Build and start the containers:
   ```bash
   docker-compose up -d
   ```

## API Endpoints

### Authentication

- `POST /auth/register` - Register a new user
- `POST /auth/login` - Login and get JWT token

### Movies

All movie endpoints require JWT authentication.

- `GET /api/movies` - Get all movies
- `GET /api/movies/:id` - Get a specific movie
- `POST /api/movies` - Create a new movie
- `PUT /api/movies/:id` - Update an existing movie
- `DELETE /api/movies/:id` - Delete a movie

## API Documentation

Swagger documentation is available at `/swagger/index.html` after starting the application.

## License

This project is licensed under the MIT License.
