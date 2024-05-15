package services

import (
    "context"
    "go-supermarket-notes-api/config"
    "go-supermarket-notes-api/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "time"
)

func CreateNote(note models.Note) error {
    collection := config.GetCollection("notes")
    note.ID = primitive.NewObjectID()
    note.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
    _, err := collection.InsertOne(context.Background(), note)
    return err
}

func GetNotes() ([]models.Note, error) {
    collection := config.GetCollection("notes")
    var notes []models.Note
    cursor, err := collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        var note models.Note
        if err = cursor.Decode(&note); err != nil {
            return nil, err
        }
        notes = append(notes, note)
    }
    return notes, nil
}

func GetNote(id string) (*models.Note, error) {
    collection := config.GetCollection("notes")
    var note models.Note
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }
    err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&note)
    return &note, err
}

func UpdateNote(id string, note models.Note) error {
    collection := config.GetCollection("notes")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    _, err = collection.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.D{
        {Key: "$set", Value: bson.M{
            "createdAt": note.CreatedAt,
            "completed": note.Completed,
            "author":    note.Author,
            "comment":   note.Comment, // Asegúrate de actualizar el campo Comment
        }},
    })
    return err
}

func DeleteNote(id string) error {
    collection := config.GetCollection("notes")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    _, err = collection.DeleteOne(context.Background(), bson.M{"_id": objID})
    return err
}
