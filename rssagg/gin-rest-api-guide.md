# Gin REST API Complete Guide

## Installation

```bash
go mod init gin-api

go get github.com/gin-gonic/gin
go get github.com/golang-jwt/jwt/v5
go get github.com/gin-contrib/cors
```

---

# Basic Server

```go
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {

    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message":"running",
        })
    })

    r.Run(":8080")
}
```

---

# API Versioning

```go
api := r.Group("/api")

v1 := api.Group("/v1")
{
    v1.GET("/users", GetUsers)
}

v2 := api.Group("/v2")
{
    v2.GET("/users", GetUsersV2)
}
```

---

# CRUD API

## GET

```go
func GetUsers(c *gin.Context) {

    c.JSON(200, users)
}
```

## GET BY ID

```go
func GetUser(c *gin.Context) {

    id := c.Param("id")

    c.JSON(200, gin.H{
        "id": id,
    })
}
```

## POST

```go
func CreateUser(c *gin.Context) {

    var user User

    c.BindJSON(&user)

    users = append(users, user)

    c.JSON(201, user)
}
```

## PUT

```go
func UpdateUser(c *gin.Context) {

}
```

## DELETE

```go
func DeleteUser(c *gin.Context) {

}
```

---

# JWT Authentication

```go
func JWTMiddleware() gin.HandlerFunc {

    return func(c *gin.Context) {

        tokenString :=
        strings.TrimPrefix(
            c.GetHeader("Authorization"),
            "Bearer ",
        )

        _, err := jwt.Parse(
            tokenString,
            func(token *jwt.Token)
            (interface{}, error) {

                return []byte(secret), nil
            },
        )

        if err != nil {

            c.JSON(401,
            gin.H{"error":"unauthorized"})

            c.Abort()

            return
        }

        c.Next()
    }
}
```

---

# File Upload

```go
func Upload(c *gin.Context) {

    file, _ := c.FormFile("file")

    c.SaveUploadedFile(
        file,
        "./uploads/"+file.Filename,
    )

    c.JSON(200, gin.H{
        "file": file.Filename,
    })
}
```

---

# Multiple Upload

```go
func UploadMany(c *gin.Context) {

    form, _ := c.MultipartForm()

    files := form.File["files"]

    for _, file := range files {

        c.SaveUploadedFile(
            file,
            "./uploads/"+file.Filename,
        )
    }
}
```

---

# Serve Static Files

```go
r.Static("/files", "./uploads")
```

---

# Video Streaming

```go
func Stream(c *gin.Context) {

    c.File("./videos/movie.mp4")
}
```

---

# Video Upload

```go
func UploadVideo(c *gin.Context) {

    video, _ :=
    c.FormFile("video")

    c.SaveUploadedFile(
        video,
        "./videos/"+video.Filename,
    )
}
```

---

# FFmpeg Video Processing

```go
cmd := exec.Command(
    "ffmpeg",
    "-i",
    "input.mp4",
    "-vf",
    "scale=1280:720",
    "output.mp4",
)

cmd.Run()
```

---

# PostgreSQL

```go
db, err :=
gorm.Open(
postgres.Open(dsn),
&gorm.Config{},
)
```

---

# Redis

```go
rdb := redis.NewClient(
&redis.Options{
    Addr: "localhost:6379",
})
```

---

# Swagger

```bash
go install github.com/swaggo/swag/cmd/swag@latest

swag init
```

---

# WebSocket

```go
func HandleWS(c *gin.Context) {

    conn, _ :=
    upgrader.Upgrade(
    c.Writer,
    c.Request,
    nil,
    )

    defer conn.Close()

    for {

        _, msg, _ :=
        conn.ReadMessage()

        conn.WriteMessage(
        websocket.TextMessage,
        msg,
        )
    }
}
```

---

# Background Jobs

```go
go func() {

    for {

        processJobs()

        time.Sleep(
        time.Second * 10,
        )
    }
}()
```

---

# Graceful Shutdown

```go
srv := &http.Server{
    Addr: ":8080",
    Handler: r,
}
```

---

# Production Stack

* Gin
* PostgreSQL
* GORM
* Redis
* JWT
* Swagger
* Docker
* Kubernetes
* Prometheus
* OpenTelemetry
* AWS S3
* FFmpeg
* Kafka
* RabbitMQ
* Nginx

```
```
