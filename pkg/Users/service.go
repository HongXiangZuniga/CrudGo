package users

type UserServices interface {
	GetUserById(id int) (*User, error)
	GetUsersByCountrys(country string) (*[]User, error)
	UpdateUser(user User) error
	CreateUser(newUser User) error
	DeleteUser(email string) error
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

func (port *port) GetUsersByCountrys(country string) (*[]User, error) {
	users, err := port.repoMongo.FindUsersByStringField("country", country)
	if err != nil {
		return nil, err
	}
	return users, nil
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
