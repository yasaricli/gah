package gah

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginStruct struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3"`
}

type RegisterStruct struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3"`
}

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
