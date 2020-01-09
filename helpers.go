package gah

import (
	"context"
	"crypto/rand"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// TokenGenerator new token generator
func TokenGenerator() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// ComparePasswords check password
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {

	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)

	if err != nil {
		return false
	}

	return true
}

// GetUserByEmail the user receives the e-mail and returns the user.
func GetUserByEmail(email string) (UserStruct, error) {
	var user UserStruct
	collection := GetCollection()
	doc := collection.FindOne(context.TODO(), bson.M{"email": email})

	err := doc.Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}

// GetUserByID The user receives _id and returns the user.
func GetUserByID(_id interface{}) (UserStruct, error) {
	var user UserStruct
	collection := GetCollection()
	doc := collection.FindOne(context.TODO(), bson.M{"_id": _id})

	err := doc.Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}

// CreateUser You can insert a new user.
func CreateUser(email string, password string) UserStruct {
	pass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	collection := GetCollection()
	newUser := &UserRegisterStruct{
		Email:     email,
		Password:  string(pass),
		CreatedAt: time.Now(),
		Tokens:    []TokenStruct{},
	}

	insertResult, _ := collection.InsertOne(context.TODO(), newUser)
	user, _ := GetUserByID(insertResult.InsertedID)

	return user
}

// InsertHashedLoginToken Add a new auth token to the user's account
func InsertHashedLoginToken(id primitive.ObjectID) string {
	collection := GetCollection()
	token := TokenGenerator()

	newToken := &TokenStruct{
		Token:     token,
		CreatedAt: time.Now(),
	}

	collection.UpdateOne(context.TODO(),
		bson.M{"_id": id},
		bson.M{
			"$addToSet": bson.M{
				"tokens": newToken,
			},
		},
	)

	return token
}
