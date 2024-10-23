package users

import (
	"os"
	"strconv"

	"github.com/HongXiangZuniga/CrudGoExample/pkg/utils"
)

type UserServices interface {
	GetUserById(id int) (*User, error)
	GetUsersByField(field, value string, page int) (*[]User, error)
	GetAllUser(page int) (*[]User, error)
	UpdateUser(user User) error
	CreateUser(newUser User) error
	DeleteUser(id int) error
	NextPageExist(id int) error
}

type port struct {
	repoMongo UsersMongoRepo
}

func NewUserServices(repoMongo UsersMongoRepo) UserServices {
	return &port{repoMongo}
}

func (port *port) GetUserById(id int) (*User, error) {
	user, err := port.repoMongo.FindUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (port *port) GetUsersByField(field, value string, page int) (*[]User, error) {
	users, err := port.repoMongo.FindUsersByStringField(field, value)
	if err != nil {
		return nil, err
	}
	pagination := os.Getenv("ELEMENTS_TO_PAGINATE")
	paginationInt, err := strconv.Atoi(pagination)
	if err != nil {
		return nil, err
	}
	userPagination := *users
	if page*paginationInt > len(userPagination) {
		return nil, utils.PageNotValid()
	}
	last := 0
	if ((page * paginationInt) + paginationInt) > len(userPagination) {
		last = len(userPagination) - page*paginationInt
		last = (page * paginationInt) + last
	} else {
		last = ((page * paginationInt) + paginationInt)
	}
	userPagination = userPagination[page*paginationInt : last]
	return &userPagination, nil
}

func (port *port) GetAllUser(page int) (*[]User, error) {
	users, err := port.repoMongo.GetAllUser()
	if err != nil {
		return nil, err
	}
	pagination := os.Getenv("ELEMENTS_TO_PAGINATE")
	paginationInt, err := strconv.Atoi(pagination)
	if err != nil {
		return nil, err
	}
	userPagination := *users
	if page*paginationInt > len(userPagination) {
		return nil, utils.PageNotValid()
	}
	last := 0
	if ((page * paginationInt) + paginationInt) > len(userPagination) {
		last = len(userPagination) - page*paginationInt
		last = (page * paginationInt) + last
	} else {
		last = ((page * paginationInt) + paginationInt)
	}
	userPagination = userPagination[page*paginationInt : last]
	return &userPagination, nil
}

func (port *port) UpdateUser(user User) error {
	return nil
}
func (port *port) CreateUser(newUser User) error {
	return port.repoMongo.CreateUser(newUser)
}
func (port *port) DeleteUser(id int) error {
	return port.repoMongo.DeleteUser(id)
}
func (port *port) NextPageExist(page int) error {
	users, err := port.repoMongo.GetAllUser()
	if err != nil {
		return err
	}
	pagination := os.Getenv("ELEMENTS_TO_PAGINATE")
	paginationInt, err := strconv.Atoi(pagination)
	if err != nil {
		return err
	}
	userPagination := *users
	if (page+1)*paginationInt > len(userPagination) {
		return utils.PageNotValid()
	}
	return nil
}
