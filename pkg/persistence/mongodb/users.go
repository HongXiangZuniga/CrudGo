package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	users "github.com/HongXiangZuniga/CrudGoExample/pkg/Users"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type storage struct {
	db *mongo.Database
}

type User = users.User

func NewUserRepo(db *mongo.Database) users.UsersMongoRepo {
	return &storage{db}
}

func (stg *storage) FindUserById(id int) (*users.User, error) {
	mongoCollection := os.Getenv("MONGO_COLLECTION")
	collection := stg.db.Collection(mongoCollection)
	ctx := context.Background()
	filter := bson.M{
		"id": bson.M{"$eq": id},
	}
	var userSearched bson.M
	err := collection.FindOne(ctx, filter).Decode(&userSearched)
	if err != nil {
		return nil, err
	}
	user := users.User{
		Id:        userSearched["id"].(int32),
		Name:      userSearched["name"].(string),
		Email:     userSearched["email"].(string),
		Age:       userSearched["age"].(int32),
		Country:   userSearched["country"].(string),
		EntryDate: userSearched["entryDate"].(primitive.DateTime),
	}
	return &user, nil
}

func (stg *storage) GetAllUser() (*[]User, error) {
	mongoCollection := os.Getenv("MONGO_COLLECTION")
	collection := stg.db.Collection(mongoCollection)
	ctx := context.Background()
	filter := bson.M{}
	result, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	Users := make([]users.User, 0)
	for result.Next(ctx) {
		var userSearched bson.M
		if err = result.Decode(&userSearched); err != nil {
			log.Fatal(err)
		} else {
			user := users.User{
				Id:        userSearched["id"].(int32),
				Name:      userSearched["name"].(string),
				Email:     userSearched["email"].(string),
				Age:       userSearched["age"].(int32),
				Country:   userSearched["country"].(string),
				EntryDate: userSearched["entryDate"].(primitive.DateTime),
			}
			Users = append(Users, user)
		}
	}
	return &Users, nil
}

func (stg *storage) FindUsersByStringField(field, value string) (*[]User, error) {
	mongoCollection := os.Getenv("MONGO_COLLECTION")
	collection := stg.db.Collection(mongoCollection)
	ctx := context.Background()
	filter := bson.M{}
	valueInt, err := strconv.Atoi(value)
	if err == nil {
		filter = bson.M{
			field: bson.M{"$eq": valueInt},
		}
	} else {
		filter = bson.M{
			field: bson.M{"$eq": value},
		}
	}
	log.Println(filter)
	result, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	Users := make([]users.User, 0)
	for result.Next(ctx) {
		var userSearched bson.M
		if err = result.Decode(&userSearched); err != nil {
			log.Fatal(err)
		} else {
			user := users.User{
				Id:        userSearched["id"].(int32),
				Name:      userSearched["name"].(string),
				Email:     userSearched["email"].(string),
				Age:       userSearched["age"].(int32),
				Country:   userSearched["country"].(string),
				EntryDate: userSearched["entryDate"].(primitive.DateTime),
			}
			Users = append(Users, user)
		}
	}
	return &Users, nil
}

func (stg *storage) CreateUser(newUser users.User) error {
	var lastUser User
	mongoCollection := os.Getenv("MONGO_COLLECTION")
	collection := stg.db.Collection(mongoCollection)
	ctx := context.Background()
	options := options.Find().SetSort(map[string]int{"id": -1}).SetLimit(1)
	result, err := collection.Find(ctx, bson.M{}, options)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	for result.Next(ctx) {
		err = result.Decode(&lastUser)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	newUser.Id = lastUser.Id + 1
	newUser.EntryDate = primitive.NewDateTimeFromTime(time.Now())
	_, err = collection.InsertOne(ctx, newUser)
	if err != nil {
		return err
	}

	return nil
}

func (stg *storage) DeleteUser(idUser int) error {
	mongoCollection := os.Getenv("MONGO_COLLECTION")
	collection := stg.db.Collection(mongoCollection)
	ctx := context.Background()
	filter := bson.M{
		"id": bson.M{"$eq": idUser},
	}
	_, err := collection.DeleteOne(ctx, filter)
	return err
}

/*
func (stg *storage) Update(user users.User) error {

	return nil
*/
