package infrastracture

import (
	"project/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDB struct {
	ID       primitive.ObjectID `bson:"_id"`
	User domain.User
}

func NewUserDB() *UserDB {
	return &UserDB{
		ID: primitive.NewObjectID(),
	}
}
