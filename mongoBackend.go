package gah

import (
	"context"
	"crypto/rand"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type MongoBackend struct {
	MongoUrl       string
	DbName         string
	CollectionName string
}

func NewMongoBackend(mongoUrl, dbName, collectionName string) *MongoBackend {
	return &MongoBackend{MongoUrl: mongoUrl, DbName: dbName, CollectionName: collectionName}
}

func (m *MongoBackend) TokenGenerator() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func (m *MongoBackend) ComparePasswords(hashedPwd string, plainPwd []byte) bool {

	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)

	if err != nil {
		return false
	}

	return true
}

func (m *MongoBackend) GetUserByEmail(email string) (UserStruct, error) {
	var user UserStruct
	collection := GetCollection(m.MongoUrl, m.DbName, m.CollectionName)
	doc := collection.FindOne(context.TODO(), bson.M{"email": email})

	err := doc.Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (m *MongoBackend) GetUserByID(_id interface{}) (UserStruct, error) {
	var user UserStruct
	collection := GetCollection(m.MongoUrl, m.DbName, m.CollectionName)
	doc := collection.FindOne(context.TODO(), bson.M{"_id": _id})

	err := doc.Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (m *MongoBackend) CreateUser(email string, password string) UserStruct {
	pass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	collection := GetCollection(m.MongoUrl, m.DbName, m.CollectionName)
	newUser := &UserRegisterStruct{
		Email:     email,
		Password:  string(pass),
		CreatedAt: time.Now(),
		Tokens:    []TokenStruct{},
	}

	insertResult, _ := collection.InsertOne(context.TODO(), newUser)
	user, _ := m.GetUserByID(insertResult.InsertedID)

	return user
}

func (m *MongoBackend) InsertHashedLoginToken(id, token string) string {
	collection := GetCollection(m.MongoUrl, m.DbName, m.CollectionName)
	token = m.TokenGenerator()

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

func (m *MongoBackend) GetUserByToken(id, token string) (UserStruct, error) {
	var user UserStruct
	collection := GetCollection(m.MongoUrl, m.DbName, m.CollectionName)
	doc := collection.FindOne(context.TODO(),
		bson.M{
			"_id":          id,
			"tokens.token": token,
		},
	)

	err := doc.Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func getClient(mongoUrl string) *mongo.Client {

	clientOptions := options.Client().ApplyURI(mongoUrl)
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func GetCollection(mongoUrl, dbName, collectionName string) *mongo.Collection {
	collection := getClient(mongoUrl).Database(dbName).Collection(collectionName)
	return collection
}
