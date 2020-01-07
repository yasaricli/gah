package gah

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// GetUserEmail the user receives the e-mail and returns the user.
func GetUserEmail(email string) (UserStruct, error) {
	var user UserStruct
	collection := GetCollection()
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
	collection := GetCollection()
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

	collection := GetCollection()
	insertResult, _ := collection.InsertOne(context.TODO(), UserRegisterStruct{
		Email:     email,
		Password:  string(pass),
		CreatedAt: time.Now(),
		Tokens:    []TokenStruct{},
	})

	user, _ := GetUserID(insertResult.InsertedID)

	return user
}
