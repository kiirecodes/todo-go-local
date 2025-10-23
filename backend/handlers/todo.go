package handlers

import (
    "context"
    "net/http"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

    "todo_backend/models"
    "todo_backend/repo"
)

func collection() *mongo.Collection {
    return repo.Client.Database(repo.DBName).Collection("todos")
}

// GetTodos supports ?q=search & ?completed=true/false & ?priority=High
func GetTodos(c *gin.Context) {
    q := c.Query("q")
    completed := c.Query("completed")
    priority := c.Query("priority")
    limitStr := c.Query("limit")
    limit := int64(100)
    if limitStr != "" {
        if v, err := strconv.ParseInt(limitStr, 10, 64); err == nil { limit = v }
    }

    filter := bson.M{}
    if q != "" {
        filter["title"] = bson.M{"$regex": q, "$options": "i"}
    }
    if completed == "true" { filter["completed"] = true }
    if completed == "false" { filter["completed"] = false }
    if priority != "" { filter["priority"] = priority }

    opts := options.Find().SetSort(bson.D{{Key: "order", Value: 1}}).SetLimit(limit)

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    cur, err := collection().Find(ctx, filter, opts)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    var todos []models.Todo
    if err := cur.All(ctx, &todos); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
    var payload models.Todo
    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    payload.CreatedAt = time.Now().UTC()
    if payload.Priority == "" { payload.Priority = "Medium" }
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    res, err := collection().InsertOne(ctx, payload)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    payload.ID = res.InsertedID.(primitive.ObjectID)
    c.JSON(http.StatusCreated, payload)
}

func UpdateTodo(c *gin.Context) {
    id := c.Param("id")
    oid, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }
    var payload models.Todo
    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    update := bson.M{"$set": bson.M{
        "title": payload.Title,
        "description": payload.Description,
        "completed": payload.Completed,
        "priority": payload.Priority,
        "due_date": payload.DueDate,
        "order": payload.Order,
    }}
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    res := collection().FindOneAndUpdate(ctx, bson.M{"_id": oid}, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
    var updated models.Todo
    if err := res.Decode(&updated); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
        return
    }
    c.JSON(http.StatusOK, updated)
}

func DeleteTodo(c *gin.Context) {
    id := c.Param("id")
    oid, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    _, err = collection().DeleteOne(ctx, bson.M{"_id": oid})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}
