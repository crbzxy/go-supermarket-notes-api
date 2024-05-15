package services

import (
    "context"
    "go-supermarket-notes-api/config"
    "go-supermarket-notes-api/models"
    "golang.org/x/crypto/bcrypt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(user models.User) error {
    // Codificar la contraseña
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    // Reemplazar la contraseña plana con la codificada
    user.Password = string(hashedPassword)

    // Insertar el usuario en la base de datos
    collection := config.GetCollection("users")
    _, err = collection.InsertOne(context.Background(), user)
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
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }
    err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
    return &user, err
}

func UpdateUser(id string, user models.User) error {
    collection := config.GetCollection("users")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    _, err = collection.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.D{
        {Key: "$set", Value: bson.M{
            "username": user.Username,
            "password": user.Password,
        }},
    })
    return err
}

func DeleteUser(id string) error {
    collection := config.GetCollection("users")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    _, err = collection.DeleteOne(context.Background(), bson.M{"_id": objID})
    return err
}

func VerifyPassword(storedPassword, providedPassword string) error {
    return bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
}

func GetUserByUsername(username string) (*models.User, error) {
    collection := config.GetCollection("users")
    var user models.User
    err := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
    return &user, err
}
