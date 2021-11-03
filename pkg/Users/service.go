package users

type UserServices interface {
	GetUser(id int) (*User, error)
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

func (port *port) GetUser(id int) (*User, error) {
	user, err := port.repoMongo.FindUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
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
