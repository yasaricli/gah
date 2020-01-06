package utils

import (
	"context"

	"../db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStruct struct {
	ID    primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Email string             `json:"email"`
}

func GetUser(email string) (UserStruct, error) {
	var user UserStruct
	collection := db.GetCollection("users")
	doc := collection.FindOne(context.TODO(), bson.M{"email": email})

	err := doc.Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}
