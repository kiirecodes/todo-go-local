package repo

import (
    "log"

    "go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client
var DBName string

func Init(client *mongo.Client, dbName string) {
    Client = client
    DBName = dbName
    if Client == nil {
        log.Fatal("mongo client is nil in repo.Init")
    }
}
