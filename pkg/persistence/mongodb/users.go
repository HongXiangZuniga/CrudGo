package mongodb

import (
	"context"
	"os"

	users "github.com/HongXiangZuniga/CrudGoExample/pkg/Users"
	"github.com/HongXiangZuniga/CrudGoExample/pkg/utils"
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

func (stg *storage) FindUser(email string) (*users.User, error) {
	mongoCollection := os.Getenv("MONGO_COLLECTION")
	collection := stg.db.Collection(mongoCollection)
	ctx := context.Background()
	filter := bson.M{
		email: bson.M{"$eq": email},
	}
	var user bson.M
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err == nil {
		//User exist
		return nil, utils.UserisExistError()
	}

	return nil, nil
}

func (stg *storage) CreateUser(user users.User) error {
	return nil
}

func (stg *storage) DeleteUser(email string) error {
	return nil
}
