package main

import (
    "context"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

    "todo_backend/handlers"
    "todo_backend/repo"
)

func main() {
    // Load environment variables
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        mongoURI = "mongodb://localhost:27017" // default for local dev
    }
    dbName := os.Getenv("MONGO_DB")
    if dbName == "" {
        dbName = "todo_db"
    }
    frontendOrigin := os.Getenv("FRONTEND_ORIGIN") // e.g. https://your-frontend.vercel.app
    if frontendOrigin == "" {
        frontendOrigin = "http://localhost:3000"
    }

    // Connect to Mongo
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatal("failed to connect to mongo:", err)
    }
    if err := client.Ping(ctx, nil); err != nil {
        log.Println("warning: could not ping mongo - ensure Mongo is running for local dev:", err)
    }

    // Initialize repository layer
    repo.Init(client, dbName)

    // Setup Gin
    router := gin.Default()

    // CORS config - allow frontend origin and common headers
    corsCfg := cors.Config{
        AllowOrigins:     []string{frontendOrigin, "http://localhost:3000"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }
    router.Use(cors.New(corsCfg))

    api := router.Group("/api")
    {
        api.GET("/health", func(c *gin.Context){ c.JSON(http.StatusOK, gin.H{"status":"ok"}) })
        api.GET("/todos", handlers.GetTodos)
        api.POST("/todos", handlers.CreateTodo)
        api.PUT("/todos/:id", handlers.UpdateTodo)
        api.DELETE("/todos/:id", handlers.DeleteTodo)
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Println("listening on :" + port)
    router.Run("0.0.0.0:" + port)
}
