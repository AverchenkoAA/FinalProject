package infrastracture

import (
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ComputerRepository interface {
	//List of functions for COMPUTER
	InsertComputer(comp *ComputerDB) error
	DeleteComputerByID(idValue string) error
	UpdateComputerByID(idValue, updateField string, updateValue interface{}) error
	FindOneComputer(filterField string, filterValue interface{}) (*ComputerDB, error)
	FindAllComputer() ([]*ComputerDB, error)
	FindComputer(filterField string, filterValue interface{}) ([]*ComputerDB, error)
	FindByID(idValue string) (*ComputerDB, error)
}

////////////////////////////////////////////////////

type compDB struct{}

func NewCompRep() ComputerRepository {
	return &compDB{}
}

func (cbd *compDB) InsertComputer(comp *ComputerDB) error {
	//Open - Close connection
	con:=NewConnect()
	ctx, cancel, client, collection, err := con.OpenConnection(CONST_collComputers)
	defer con.CloseConnection(ctx, cancel, client)
	/////////////////////////
	_, err = collection.InsertOne(*ctx, comp)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (cbd *compDB) DeleteComputerByID(idValue string) error {
	//Open - Close connection
	con:=NewConnect()
	ctx, cancel, client, collection, err := con.OpenConnection(CONST_collComputers)
	defer con.CloseConnection(ctx, cancel, client)
	/////////////////////////
	id,err:=primitive.ObjectIDFromHex(idValue)
	filter := bson.D{{"_id", id}}
	_, err = collection.DeleteOne(*ctx, filter)

	return err
}

func (cbd *compDB) UpdateComputerByID(idValue, updateField string, updateValue interface{}) error {
	//Open - Close connection
	con:=NewConnect()
	ctx, cancel, client, collection, err := con.OpenConnection(CONST_collComputers)
	defer con.CloseConnection(ctx, cancel, client)
	/////////////////////////
	id,err:=primitive.ObjectIDFromHex(idValue)
	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{
			{updateField, updateValue},
		}},
	}
	_, err = collection.UpdateOne(*ctx, filter, update)

	return err
}

func (cbd *compDB) FindOneComputer(filterField string, filterValue interface{}) (*ComputerDB, error) {

	filter := bson.D{{filterField, filterValue}}
	var result *ComputerDB
	//Open - Close connection
	con:=NewConnect()
	ctx, cancel, client, collection, err := con.OpenConnection(CONST_collComputers)
	defer con.CloseConnection(ctx, cancel, client)
	/////////////////////////

	err = collection.FindOne(*ctx, filter).Decode(&result)

	return result, err
}

func (cbd *compDB) FindAllComputer() ([]*ComputerDB, error) {
	//Open - Close connection
	con:=NewConnect()
	ctx, cancel, client, collection, err := con.OpenConnection(CONST_collComputers)
	defer con.CloseConnection(ctx, cancel, client)
	/////////////////////////

	filter := bson.M{}
	var results []*ComputerDB
	var cur *mongo.Cursor
	cur, err = collection.Find(*ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(*ctx) {
		var c ComputerDB
		err = cur.Decode(&c)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &c)
	}
	if err = cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(*ctx)

	return results, err
}

func (cbd *compDB) FindComputer(filterField string, filterValue interface{}) ([]*ComputerDB, error) {
	//Open - Close connection
	con:=NewConnect()
	ctx, cancel, client, collection, err := con.OpenConnection(CONST_collComputers)
	defer con.CloseConnection(ctx, cancel, client)
	/////////////////////////

	var results []*ComputerDB
	var cur *mongo.Cursor

	cur, err = collection.Find(*ctx, bson.M{filterField: filterValue})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(*ctx) {
		var c ComputerDB
		err = cur.Decode(&c)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &c)
	}
	if err = cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(*ctx)

	return results, err
}

func (cbd *compDB) FindByID(idValue string) (*ComputerDB, error) {
	
	var result *ComputerDB
	//Open - Close connection
	con:=NewConnect()
	ctx, cancel, client, collection, err := con.OpenConnection(CONST_collComputers)
	defer con.CloseConnection(ctx, cancel, client)
	/////////////////////////
	id,err:=primitive.ObjectIDFromHex(idValue)
	filter := bson.D{{"_id", id}}
	err = collection.FindOne(*ctx, filter).Decode(&result)

	return result, err
}
