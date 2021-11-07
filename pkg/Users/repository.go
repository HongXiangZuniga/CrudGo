package users

type UsersMongoRepo interface {
	FindUserById(id int) (*User, error)
	FindUsersByStringField(field, value string) (*[]User, error)
}
