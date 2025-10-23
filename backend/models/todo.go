package models

import "time"
import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Title       string             `bson:"title" json:"title"`
    Description string             `bson:"description,omitempty" json:"description,omitempty"`
    Completed   bool               `bson:"completed" json:"completed"`
    Priority    string             `bson:"priority,omitempty" json:"priority,omitempty"`
    DueDate     *time.Time         `bson:"due_date,omitempty" json:"due_date,omitempty"`
    CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
    Order       int                `bson:"order" json:"order"`
}
