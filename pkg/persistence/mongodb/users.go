package mongodb

import (
	"context"
	"os"

	users "github.com/HongXiangZuniga/CrudGoExample/pkg/Users"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type storage struct {
	db *mongo.Database
}

type User = users.User

func NewUserRepo(db *mongo.Database) users.UsersMongoRepo {
	return &storage{db}
}

func (stg *storage) FindUser(id int) (*users.User, error) {
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

func (stg *storage) CreateUser(user users.User) error {
	mongoCollection := os.Getenv("MONGO_COLLECTION")
	collection := stg.db.Collection(mongoCollection)
	ctx := context.Background()
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (stg *storage) DeleteUser(email string) error {
	mongoCollection := os.Getenv("MONGO_COLLECTION")
	collection := stg.db.Collection(mongoCollection)
	ctx := context.Background()
	filter := bson.M{
		email: bson.M{"$eq": email},
	}
	_, err := collection.DeleteOne(ctx, filter)
	return err
}

func (stg *storage) Update(user users.User) error {

	return nil
}
