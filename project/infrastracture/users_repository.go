package infrastracture

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"fmt"
)

type UserRepository interface {
	
	//List of functions for USER
	InsertUser(user *UserDB) error
	DeleteUserByID(idValue string) error
	UpdateUserByID(idValue, updateField string, updateValue interface{}) error
	FindOneUser(filterField string, filterValue interface{}) (*UserDB, error)
	FindAllUser() ([]*UserDB, error)
	FindUser(filterField string, filterValue interface{}) ([]*UserDB, error)
	FindByID(idValue string) (*UserDB, error)
}


type userDB struct{}

func NewUserRep() UserRepository {
	return &userDB{}
}
//Create (if it need) default ADMIN's user
func init(){
	userRep:=NewUserRep()
	_,err:=userRep.FindOneUser("user.login","ADMIN")
	if err!=nil{
		fmt.Println("Error while finding ADMIN")
		fmt.Println("Creating default ADMIN")
		newUser:=NewUserDB()
		newUser.User.Login="ADMIN"
		newUser.User.Password="25d55ad283aa400af464c76d713c07ad"
		newUser.User.UserRights="admin"
		err=userRep.InsertUser(newUser)
		if err!=nil{
			fmt.Println("Error while insert default ADMIN")
			return
		}
		return
	}
}

func (udb *userDB) InsertUser(user *UserDB) error {
	//Open - Close connection
	con:=NewConnect()
	ctx, cancel, client, collection, err := con.OpenConnection(CONST_collUsers)
	defer con.CloseConnection(ctx, cancel, client)
	/////////////////////////
	_, err = collection.InsertOne(*ctx, user)
	return err
}

func (udb *userDB) DeleteUserByID(idValue string) error {
	//Open - Close connection
	con:=NewConnect()
	ctx, cancel, client, collection, err := con.OpenConnection(CONST_collUsers)
	defer con.CloseConnection(ctx, cancel, client)
	/////////////////////////
	id,err:=primitive.ObjectIDFromHex(idValue)
	filter := bson.D{{"_id", id}}
	_, err = collection.DeleteOne(*ctx, filter)

	return err
}

func (udb *userDB) UpdateUserByID(idValue, updateField string, updateValue interface{}) error {
	//Open - Close connection
	con:=NewConnect()
	ctx, cancel, client, collection, err := con.OpenConnection(CONST_collUsers)
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

func (udb *userDB) FindOneUser(filterField string, filterValue interface{}) (*UserDB, error) {

	filter := bson.D{{filterField, filterValue}}
	var result *UserDB
	//Open - Close connection
	con:=NewConnect()
	ctx, cancel, client, collection, err := con.OpenConnection(CONST_collUsers)
	defer con.CloseConnection(ctx, cancel, client)
	/////////////////////////

	err = collection.FindOne(*ctx, filter).Decode(&result)

	return result, err
}

func (udb *userDB) FindAllUser() ([]*UserDB, error) {
	//Open - Close connection
	con:=NewConnect()
	ctx, cancel, client, collection, err := con.OpenConnection(CONST_collUsers)
	defer con.CloseConnection(ctx, cancel, client)
	/////////////////////////

	filter := bson.M{}
	var results []*UserDB
	var cur *mongo.Cursor
	cur, err = collection.Find(*ctx, filter)


	for cur.Next(*ctx) {
		var c UserDB
		err = cur.Decode(&c)

		results = append(results, &c)
	}

	cur.Close(*ctx)

	return results, err
}

func (udb *userDB) FindUser(filterField string, filterValue interface{}) ([]*UserDB, error) {
	//Open - Close connection
	con:=NewConnect()
	ctx, cancel, client, collection, err := con.OpenConnection(CONST_collUsers)
	defer con.CloseConnection(ctx, cancel, client)
	/////////////////////////

	var results []*UserDB
	var cur *mongo.Cursor

	cur, err = collection.Find(*ctx, bson.M{filterField: filterValue})

	for cur.Next(*ctx) {
		var c UserDB
		err = cur.Decode(&c)

		results = append(results, &c)
	}

	cur.Close(*ctx)

	return results, err
}

func (udb *userDB) FindByID(idValue string) (*UserDB, error) {
	
	var result *UserDB
	//Open - Close connection
	con:=NewConnect()
	ctx, cancel, client, collection, err := con.OpenConnection(CONST_collUsers)
	defer con.CloseConnection(ctx, cancel, client)
	/////////////////////////
	id,err:=primitive.ObjectIDFromHex(idValue)
	filter := bson.D{{"_id", id}}
	err = collection.FindOne(*ctx, filter).Decode(&result)

	return result, err
}
