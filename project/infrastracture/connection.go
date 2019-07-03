package infrastracture

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
 
)

type Connection interface {
	OpenConnection(collectionName string) (*context.Context, context.CancelFunc, *mongo.Client, *mongo.Collection, error)
	CloseConnection(ctx *context.Context, cancel context.CancelFunc, client *mongo.Client) error
}

//Connection constants//////////////////////////////
const (
	CONST_dbConnect string = "mongodb://127.0.0.1:27017"
	CONST_dbName string = "pcBase"
	CONST_collComputers string = "Computers"
	CONST_collUsers string = "Users"
)


////////////////////////////////////////////////////

type connect struct{}

func NewConnect() Connection {
	return &connect{}
}

//Connection block to data base/////////////////////////////////////////////
///////Open connection////////////////////////////////////////////////////////
func (udb *connect) OpenConnection(collectionName string) (*context.Context, context.CancelFunc, *mongo.Client, *mongo.Collection, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.NewClient(options.Client().ApplyURI(CONST_dbConnect))

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(CONST_dbName).Collection(collectionName)
	return &ctx, cancel, client, collection, err
}

///////Close connection////////////////////////////////////////////////////////
func (udb *connect) CloseConnection(ctx *context.Context, cancel context.CancelFunc, client *mongo.Client) error {
	err := client.Disconnect(*ctx)
	if err != nil {
		log.Fatal(err)
	}
	cancel()
	return err
}

////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////