package config

import (
    "context"
    "log"
    "os"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() {
    clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))

    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    DB = client
    log.Println("Connected to MongoDB!")
}

func GetCollection(collectionName string) *mongo.Collection {
    return DB.Database(os.Getenv("DB_NAME")).Collection(collectionName)
}
