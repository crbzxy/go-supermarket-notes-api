package services

import (
    "context"
    "go-supermarket-notes-api/config"
    "go-supermarket-notes-api/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user models.User) error {
    collection := config.GetCollection("users")
    _, err := collection.InsertOne(context.Background(), user)
    return err
}

func GetUsers() ([]models.User, error) {
    collection := config.GetCollection("users")
    var users []models.User
    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        var user models.User
        if err = cursor.Decode(&user); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

func GetUser(id string) (*models.User, error) {
    collection := config.GetCollection("users")
    var user models.User
    objID, _ := primitive.ObjectIDFromHex(id)
    err := collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
    return &user, err
}

func UpdateUser(id string, user models.User) error {
    collection := config.GetCollection("users")
    objID, _ := primitive.ObjectIDFromHex(id)
    _, err := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.D{
        {Key: "$set", Value: bson.M{
            "username": user.Username,
            "password": user.Password,
        }},
    })
    return err
}

func DeleteUser(id string) error {
    collection := config.GetCollection("users")
    objID, _ := primitive.ObjectIDFromHex(id)
    _, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})
    return err
}
