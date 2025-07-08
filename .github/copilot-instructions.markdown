# 🧪 Production-Ready Go Backend with Gin, GORM, and PostgreSQL

This guide explains how to build a **scalable and maintainable Go backend** using **Gin**, **GORM**, and **PostgreSQL** — including folder structure, routing, migrations, and best practices.

---

## 📁 Folder Structure

```
/cmd/                  → Application entrypoints
/internal/
  config/              → Environment configs and setup
  db/                  → DB connection and migrations
  model/               → GORM models
  repository/          → Data access layer
  service/             → Business logic
  handler/             → HTTP handlers/controllers
  middleware/          → Custom Gin middleware
  route/               → Route definitions and grouping
/pkg/                  → Reusable packages (e.g., logger, utils)
```

**Highlights:**
- `internal/` protects from external imports.
- Separation of concerns: handler ≠ service ≠ repository.
- Easy to test, scale, and refactor.

---

## 🌐 Routing

```go
func RegisterAPIRoutes(r *gin.Engine, h *handler.Handler) {
    v1 := r.Group("/api/v1")
    {
        user := v1.Group("/users")
        user.GET("/", h.ListUsers)
        user.POST("/", h.CreateUser)
    }
}
```

**Best Practices:**
- Use API versioning from the start (`/v1`, `/v2`).
- Group routes by feature/resource.
- Use middleware per route group if needed.

---

## 🧩 Dependency Injection

```go
type Handler struct {
    UserService *service.UserService
}

func NewHandler(userService *service.UserService) *Handler {
    return &Handler{UserService: userService}
}
```

- Keep wiring in `main.go` or initializer packages.
- Helps test services and handlers separately.

---

## 🧱 GORM & PostgreSQL Setup

### Connection

Use a `Config` struct from env vars to create a DSN:

```go
dsn := fmt.Sprintf(
  "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
  cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.SSLMode,
)
```

### Connection Pooling

```go
sqlDB.SetMaxOpenConns(20)
sqlDB.SetMaxIdleConns(10)
sqlDB.SetConnMaxLifetime(time.Hour)
```

### Model Design

```go
type User struct {
    ID        uint      `gorm:"primaryKey"`
    Email     string    `gorm:"uniqueIndex;not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

**Tips:**
- Avoid using `gorm.Model` directly — define your own base model.
- Consider UUIDs for distributed systems.
- Use `gorm.DeletedAt` for soft deletes.

---

## 🛠 Migrations

### ⚠️ Don’t use `AutoMigrate()` in production

Instead, use a **real migration tool**:

- [golang-migrate](https://github.com/golang-migrate/migrate)
- [goose](https://github.com/pressly/goose)

### Best Practices

- Store SQL files under `/migrations/`
- Use timestamped filenames
- Run via CI/CD or `docker exec` jobs

```bash
migrate -path ./migrations -database "$DB_URL" up
```

---

## 🔐 Env Configuration

Use `github.com/caarlos0/env` for parsing:

```go
type Config struct {
    DBHost string `env:"DB_HOST"`
    DBPort string `env:"DB_PORT"`
    ...
}
```

Use `.env` files for local dev with `github.com/joho/godotenv`.

---

## 🔍 Logging

Use structured logging:
- `zap` (Uber's high-perf logger)
- `zerolog` (lightweight, zero-allocation)

Log context (user ID, request ID, errors) for every request.

---

## 🧪 Testing Strategy

### Unit Tests
- Use mocks in `service` and `repository`
- `stretchr/testify` for assertions

### Handler Tests
- `httptest.NewRecorder()` for HTTP testing

### DB Tests
- Use Docker + Postgres or SQLite in memory
- Isolate test database

---

## 📦 Deployability

### Docker

Multi-stage build:

```Dockerfile
FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go build -o server

FROM alpine
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]
```

### Process Management

- Use behind **NGINX** or **Caddy** for reverse proxy
- Run with `systemd`, Docker Compose, or Kubernetes

---

## 🚀 Production Tips

- Use `/ping` or `/healthz` endpoints
- Protect login routes with rate limiting
- Always use `GIN_MODE=release`
- Handle panics with `gin.Recovery()`
- Apply secure headers with `gin-contrib/secure`
- Monitor DB pool and memory usage
