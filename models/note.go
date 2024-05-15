package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
    ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    CreatedAt   primitive.DateTime `json:"createdAt" bson:"createdAt"`
    Completed   bool               `json:"completed" bson:"completed"`
    Author      string             `json:"author" bson:"author"`
    Comment     string             `json:"comment" bson:"comment"`
}
