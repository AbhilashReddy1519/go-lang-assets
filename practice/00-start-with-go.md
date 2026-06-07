# Production-Ready Go Backend Setup (2026)

## 1. Install Go

### Linux (Ubuntu/Debian)

```bash
sudo apt update
sudo apt install golang-go -y

go version
```

### macOS

```bash
brew install go

go version
```

### Windows

Download and install:

https://go.dev/dl/

Verify:

```powershell
go version
```

---

# 2. Project Structure

Use a clean architecture style.

```text
go-backend/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── internal/
│   ├── handlers/
│   ├── services/
│   ├── repository/
│   ├── middleware/
│   ├── config/
│   └── models/
│
├── pkg/
│
├── migrations/
│
├── docs/
│
├── scripts/
│
├── tests/
│
├── .env
├── .env.example
├── Dockerfile
├── docker-compose.yml
├── Makefile
├── go.mod
└── go.sum
```

---

# 3. Create Project

```bash
mkdir go-backend
cd go-backend

go mod init github.com/yourname/go-backend
```

---

# 4. Install Essential Packages

## HTTP Router

```bash
go get github.com/gin-gonic/gin
```

## PostgreSQL Driver

```bash
go get gorm.io/gorm
go get gorm.io/driver/postgres
```

## Environment Variables

```bash
go get github.com/joho/godotenv
```

## Logging

```bash
go get go.uber.org/zap
```

## Validation

```bash
go get github.com/go-playground/validator/v10
```

## JWT Authentication

```bash
go get github.com/golang-jwt/jwt/v5
```

## Swagger

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

---

# 5. Create Main Entry

cmd/server/main.go

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.Run(":8080")
}
```

Run:

```bash
go run cmd/server/main.go
```

Visit:

```text
http://localhost:8080/health
```

---

# 6. Setup Environment Variables

.env

```env
APP_ENV=development

PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=mydb

JWT_SECRET=super-secret-key
```

Load:

```go
godotenv.Load()
```

---

# 7. PostgreSQL with Docker

docker-compose.yml

```yaml
version: "3.9"

services:
  postgres:
    image: postgres:17

    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: password
      POSTGRES_DB: mydb

    ports:
      - "5432:5432"

    volumes:
      - postgres_data:/var/lib/postgresql/data

volumes:
  postgres_data:
```

Run:

```bash
docker compose up -d
```

Check:

```bash
docker ps
```

---

# 8. Database Connection

internal/config/database.go

```go
package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
```

---

# 9. Logging

internal/config/logger.go

```go
package config

import "go.uber.org/zap"

func NewLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}
```

Usage:

```go
logger.Info("server started")
```

---

# 10. Middleware

Example Request Logger

```go
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()

		c.Next()

		log.Printf(
			"%s %s %v",
			c.Request.Method,
			c.Request.URL.Path,
			time.Since(start),
		)
	}
}
```

---

# 11. JWT Authentication

Generate Token

```go
token := jwt.NewWithClaims(
	jwt.SigningMethodHS256,
	jwt.MapClaims{
		"user_id": 1,
	},
)

signed, err := token.SignedString(
	[]byte(os.Getenv("JWT_SECRET")),
)
```

Verify Token

```go
jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET")), nil
})
```

---

# 12. API Versioning

```go
v1 := router.Group("/api/v1")

{
	v1.GET("/users", getUsers)
	v1.POST("/users", createUser)
}
```

---

# 13. Configuration Management

Create:

```text
internal/config/
```

Store:

```go
type Config struct {
	Port string
	DBHost string
	DBUser string
}
```

Load once at startup.

Avoid using:

```go
os.Getenv(...)
```

throughout the application.

---

# 14. Migrations

Install:

```bash
go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Create migration:

```bash
migrate create -ext sql -dir migrations create_users
```

Apply:

```bash
migrate -path migrations \
-database postgres://postgres:password@localhost:5432/mydb?sslmode=disable up
```

---

# 15. Swagger Docs

Generate:

```bash
swag init -g cmd/server/main.go
```

Access:

```text
http://localhost:8080/swagger/index.html
```

---

# 16. Unit Testing

Example

```go
func TestHealth(t *testing.T) {

	req := httptest.NewRequest(
		http.MethodGet,
		"/health",
		nil,
	)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}
}
```

Run:

```bash
go test ./...
```

---

# 17. Dockerfile

```dockerfile
FROM golang:1.25 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux \
go build -o server ./cmd/server

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
```

Build:

```bash
docker build -t go-backend .
```

Run:

```bash
docker run -p 8080:8080 go-backend
```

---

# 18. Makefile

```makefile
run:
	go run cmd/server/main.go

test:
	go test ./...

build:
	go build -o server ./cmd/server

docker:
	docker build -t go-backend .
```

Usage:

```bash
make run
make test
make build
```

---

# 19. Production Deployment

Build Binary

```bash
go build -ldflags="-s -w" -o server ./cmd/server
```

Run with systemd

```bash
sudo nano /etc/systemd/system/app.service
```

```ini
[Unit]
Description=Go Backend

[Service]
ExecStart=/home/ubuntu/server
Restart=always

[Install]
WantedBy=multi-user.target
```

```bash
sudo systemctl daemon-reload
sudo systemctl enable app
sudo systemctl start app
```

---

# 20. Reverse Proxy (Nginx)

```nginx
server {
    listen 80;

    server_name api.example.com;

    location / {
        proxy_pass http://localhost:8080;

        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

Restart:

```bash
sudo systemctl restart nginx
```

---

# 21. Production Checklist

✅ Gin Framework

✅ PostgreSQL

✅ Docker

✅ JWT Authentication

✅ Validation

✅ Structured Logging

✅ Database Migrations

✅ Swagger Documentation

✅ Unit Tests

✅ Nginx Reverse Proxy

✅ Systemd Service

✅ Environment Variables

✅ API Versioning

✅ CI/CD Ready Structure

✅ Clean Architecture

---

# Recommended Backend Stack (2026)

For most startups and production systems:

```text
Go 1.25+
Gin
PostgreSQL
GORM
JWT
Docker
Nginx
Swagger
Zap Logger
golang-migrate
GitHub Actions
AWS / GCP / DigitalOcean
```

This setup is sufficient for building production-grade APIs, SaaS products, authentication systems, payment backends, microservices, and high-performance backend services.
