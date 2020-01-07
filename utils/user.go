package utils

import (
	"context"
	"log"
	"time"

	"../db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserStruct struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Email     string             `json:"email" bson:"email"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
}

type TokenStruct struct {
	Token     string
	CreatedAt time.Time
}

type UserRegisterStruct struct {
	Email     string        `json:"email" bson:"email"`
	Password  string        `json:"password" bson:"password"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt"`
	Tokens    []TokenStruct `json: "tokens" bson:"tokens"`
}

// GetUserEmail the user receives the e-mail and returns the user.
func GetUserEmail(email string) (UserStruct, error) {
	var user UserStruct
	collection := db.GetCollection()
	doc := collection.FindOne(context.TODO(), bson.M{"email": email})

	err := doc.Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}

// GetUserID The user receives _id and returns the user.
func GetUserID(_id interface{}) (UserStruct, error) {
	var user UserStruct
	collection := db.GetCollection()
	doc := collection.FindOne(context.TODO(), bson.M{"_id": _id})

	err := doc.Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}

// InsertUser You can insert a new user.
func InsertUser(email string, password string) UserStruct {
	pass, passwordError := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if passwordError != nil {
		log.Fatalln("Error on inserting new User", passwordError)
	}

	collection := db.GetCollection()
	insertResult, _ := collection.InsertOne(context.TODO(), UserRegisterStruct{
		Email:     email,
		Password:  string(pass),
		CreatedAt: time.Now(),
		Tokens:    []TokenStruct{},
	})

	user, _ := GetUserID(insertResult.InsertedID)

	return user
}
