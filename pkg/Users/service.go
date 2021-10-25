package users

import "go.uber.org/zap"

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
	return nil, nil
}
func (port *port) UpdateUser(user User) error {
	return nil
}
func (port *port) CreateUser(newUser User) error {
	return nil
}
func (port *port) DeleteUser(email string) error {
	return nil
}
