# REST API with Go, Gin, and GORM

A simple REST API built with Go, Gin framework, and GORM for user management. Supports both SQLite (for development) and PostgreSQL (for production) with optional Redis caching.

## Features

- CRUD operations for users
- SQLite for local development
- PostgreSQL for production
- Redis caching (production)
- Environment-based configuration
- JSON validation
- Automatic database migrations

## Prerequisites

- Go 1.21 or higher
- SQLite (for development)
- PostgreSQL (for production)
- Redis (for production with caching)

## Project Structure

```
restapi/
├── config/         # Configuration management
├── handlers/       # HTTP handlers
├── middleware/     # Middleware (cache, etc.)
├── models/         # Database models
├── data/          # SQLite database files (dev)
├── .env           # Production environment variables
├── .env.dev       # Development environment variables
├── go.mod         # Go module file
└── main.go        # Application entry point
```

## Setup

### Development (SQLite)

1. Clone the repository:
```bash
git clone <repository-url>
cd restapi
```

2. Install dependencies:
```bash
go mod tidy
```

3. Create development environment file:
```bash
cp .env.dev.example .env.dev
```

4. Run the application:
```bash
go run main.go
```

### Production (PostgreSQL + Redis)

1. Create production environment file:
```bash
cp .env.example .env
```

2. Configure your `.env` file:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=restapi
DB_SSL_MODE=disable

REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

SERVER_PORT=8080
```

3. Run with production config:
```bash
go run main.go
```

### API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/docs/index.html` | Swagger UI documentation |
| POST | `/api/v1/users` | Create a new user |
| GET | `/api/v1/users` | Get all users |
| GET | `/api/v1/users/:id` | Get user by ID |
| PUT | `/api/v1/users/:id` | Update user |
| DELETE | `/api/v1/users/:id` | Delete user |

### Request/Response Examples

#### Create User
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "age": 30
  }'
```

#### Get All Users
```bash
curl http://localhost:8080/api/v1/users
```

#### Get User by ID
```bash
curl http://localhost:8080/api/v1/users/1
```

#### Update User
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Doe",
    "email": "jane@example.com",
    "age": 25
  }'
```

#### Delete User
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

## Database Schema

### User Model
```go
type User struct {
    ID        uint           `json:"id" gorm:"primarykey"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
    Name      string         `json:"name" gorm:"not null"`
    Email     string         `json:"email" gorm:"uniqueIndex;not null"`
    Age       int            `json:"age"`
}
```

## Configuration

The application uses environment variables for configuration:

### Development (.env.dev)
```env
DB_TYPE=sqlite
DB_PATH=./data/dev.db
SERVER_PORT=8080
```

### Production (.env)
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=restapi
DB_SSL_MODE=disable

REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

SERVER_PORT=8080
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This project is licensed under the MIT License.
