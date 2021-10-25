package users

import (
	"github.com/HongXiangZuniga/CrudGoExample/pkg/utils"
	"go.uber.org/zap"
)

type UserServices interface {
	GetUser(email string) (*User, error)
	UpdateUser(user User) error
	CreateUser(newUser User) error
	DeleteUser(email string) error
}

type port struct {
	logger    *zap.Logger
	repoMongo UsersMongoRepo
}

func NewUserServices(logger *zap.Logger, repoMongo UsersMongoRepo) UserServices {
	return &port{logger, repoMongo}
}

func (port *port) GetUser(email string) (*User, error) {
	user, err := port.repoMongo.FindUser(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (port *port) UpdateUser(user User) error {
	return nil
}
func (port *port) CreateUser(newUser User) error {
	_, err := port.repoMongo.FindUser(newUser.Email)
	if err == nil {
		/*User is exist*/
		return utils.UserisExistError()
	}
	err = port.repoMongo.CreateUser(newUser)
	return err
}
func (port *port) DeleteUser(email string) error {
	err := port.repoMongo.DeleteUser(email)
	return err
}
