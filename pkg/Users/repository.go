package users

type UsersMongoRepo interface {
	GetAllUser() (*[]User, error)
	FindUserById(id int) (*User, error)
	FindUsersByStringField(field, value string) (*[]User, error)
	DeleteUser(idUser int) error
	CreateUser(newUser User) error
}
