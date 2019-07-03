package infrastracture

import (
	"project/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ComputerDB struct {
	ID       primitive.ObjectID `bson:"_id"`
	Computer domain.Computer
}

func NewComputerDB() *ComputerDB {
	return &ComputerDB{
		ID: primitive.NewObjectID(),
	}
}
