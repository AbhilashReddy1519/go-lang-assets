# Chi REST API Complete Guide

## Project Structure

```text
chi-api/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   ├── services/
│   └── repository/
├── uploads/
├── videos/
├── configs/
├── go.mod
└── .env
```

---

# Installation

```bash
go mod init chi-api

go get github.com/go-chi/chi/v5
go get github.com/go-chi/cors
go get github.com/go-chi/render
go get github.com/golang-jwt/jwt/v5
go get github.com/joho/godotenv
```

---

# Basic Server

```go
package main

import (
    "log"
    "net/http"

    "github.com/go-chi/chi/v5"
)

func main() {
    r := chi.NewRouter()

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("API Running"))
    })

    log.Fatal(http.ListenAndServe(":8080", r))
}
```

---

# API Versioning

```go
r.Route("/api", func(r chi.Router) {

    r.Route("/v1", func(r chi.Router) {
        r.Get("/users", GetUsersV1)
    })

    r.Route("/v2", func(r chi.Router) {
        r.Get("/users", GetUsersV2)
    })
})
```

---

# Middleware

```go
func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        log.Println(r.Method, r.URL.Path)

        next.ServeHTTP(w, r)
    })
}
```

```go
r.Use(Logger)
```

---

# CRUD REST API

## Model

```go
type User struct {
    ID uint `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}
```

## GET

```go
func GetUsers(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(users)
}
```

## GET BY ID

```go
func GetUser(w http.ResponseWriter, r *http.Request) {

    id := chi.URLParam(r, "id")

    json.NewEncoder(w).Encode(id)
}
```

## POST

```go
func CreateUser(w http.ResponseWriter, r *http.Request) {

    var user User

    json.NewDecoder(r.Body).Decode(&user)

    users = append(users, user)

    json.NewEncoder(w).Encode(user)
}
```

## PUT

```go
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}
```

## DELETE

```go
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
```

---

# JWT Authentication

## Generate Token

```go
token := jwt.NewWithClaims(jwt.SigningMethodHS256,
jwt.MapClaims{
    "user_id": 1,
})

tokenString, _ := token.SignedString([]byte(secret))
```

## Middleware

```go
func JWTMiddleware(next http.Handler) http.Handler {

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        tokenString := strings.TrimPrefix(
            r.Header.Get("Authorization"),
            "Bearer ",
        )

        _, err := jwt.Parse(tokenString,
        func(token *jwt.Token) (interface{}, error) {
            return []byte(secret), nil
        })

        if err != nil {
            http.Error(w, "Unauthorized", 401)
            return
        }

        next.ServeHTTP(w, r)
    })
}
```

---

# File Upload

```go
func UploadFile(w http.ResponseWriter, r *http.Request) {

    file, header, err :=
        r.FormFile("file")

    if err != nil {
        return
    }

    defer file.Close()

    dst, _ := os.Create(
        "./uploads/" + header.Filename,
    )

    defer dst.Close()

    io.Copy(dst, file)

    json.NewEncoder(w).Encode(map[string]string{
        "file": header.Filename,
    })
}
```

Route

```go
r.Post("/upload", UploadFile)
```

---

# Multiple File Upload

```go
func UploadMultiple(w http.ResponseWriter, r *http.Request) {

    r.ParseMultipartForm(32 << 20)

    files := r.MultipartForm.File["files"]

    for _, fileHeader := range files {

        file, _ := fileHeader.Open()

        dst, _ := os.Create(
            "./uploads/" + fileHeader.Filename,
        )

        io.Copy(dst, file)

        file.Close()
        dst.Close()
    }
}
```

---

# Serve Uploaded Files

```go
fs := http.FileServer(
    http.Dir("./uploads"),
)

r.Handle("/files/*",
http.StripPrefix("/files", fs))
```

---

# Video Streaming

```go
func StreamVideo(w http.ResponseWriter, r *http.Request) {

    video := "./videos/sample.mp4"

    http.ServeFile(w, r, video)
}
```

```go
r.Get("/video", StreamVideo)
```

---

# Video Upload

```go
func UploadVideo(w http.ResponseWriter, r *http.Request) {

    file, header, _ := r.FormFile("video")

    defer file.Close()

    dst, _ := os.Create(
        "./videos/" + header.Filename,
    )

    io.Copy(dst, file)

    dst.Close()
}
```

---

# Database Connection (Postgres)

```go
db, err := sql.Open(
    "postgres",
    connString,
)
```

---

# Pagination

```go
page := r.URL.Query().Get("page")
limit := r.URL.Query().Get("limit")
```

---

# Search

```go
search := r.URL.Query().Get("search")
```

---

# WebSocket

```go
var upgrader = websocket.Upgrader{}

func WS(w http.ResponseWriter, r *http.Request) {

    conn, _ :=
        upgrader.Upgrade(w, r, nil)

    defer conn.Close()

    for {

        _, msg, _ := conn.ReadMessage()

        conn.WriteMessage(
            websocket.TextMessage,
            msg,
        )
    }
}
```

---

# Graceful Shutdown

```go
srv := &http.Server{
    Addr: ":8080",
    Handler: r,
}

go srv.ListenAndServe()

ctx, cancel :=
context.WithTimeout(
context.Background(),
5*time.Second,
)

defer cancel()

srv.Shutdown(ctx)
```

---

# Production Recommendations

* Chi Router
* PostgreSQL
* Redis
* JWT
* Docker
* Prometheus
* OpenTelemetry
* Nginx
* AWS S3 for uploads
* FFmpeg for video processing
* Kafka/NATS for events
* Swagger/OpenAPI
* Kubernetes

```
```
