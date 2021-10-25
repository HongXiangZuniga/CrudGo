package mongodb

import (
	users "github.com/HongXiangZuniga/CrudGoExample/pkg/Users"
	"go.mongodb.org/mongo-driver/mongo"
)

type storage struct {
	db *mongo.Database
}

func NewUserRepo(db *mongo.Database) users.UsersMongoRepo {
	return &storage{db}
}

func (stg *storage) FindUser(email string) (*users.User, error) {
	return nil, nil
}

func (stg *storage) CreateUser(user users.User) error {
	return nil
}

func (stg *storage) DeleteUser(email string) error {
	return nil
}
