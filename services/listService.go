package services

import (
    "context"
    "go-supermarket-notes-api/config"
    "go-supermarket-notes-api/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateList(list models.List) error {
    collection := config.GetCollection("lists")
    _, err := collection.InsertOne(context.Background(), list)
    return err
}

func GetLists() ([]models.List, error) {
    collection := config.GetCollection("lists")
    var lists []models.List
    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        var list models.List
        if err = cursor.Decode(&list); err != nil {
            return nil, err
        }
        lists = append(lists, list)
    }
    return lists, nil
}

func GetList(id string) (*models.List, error) {
    collection := config.GetCollection("lists")
    var list models.List
    objID, _ := primitive.ObjectIDFromHex(id)
    err := collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&list)
    return &list, err
}

func UpdateList(id string, list models.List) error {
    collection := config.GetCollection("lists")
    objID, _ := primitive.ObjectIDFromHex(id)
    _, err := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.D{
        {Key: "$set", Value: bson.M{
            "name":     list.Name,
            "category": list.Category,
        }},
    })
    return err
}

func DeleteList(id string) error {
    collection := config.GetCollection("lists")
    objID, _ := primitive.ObjectIDFromHex(id)
    _, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})
    return err
}
