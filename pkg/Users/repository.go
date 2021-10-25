package users

type UsersMongoRepo interface {
	FindUser(email string) (*User, error)
	CreateUser(user User) error
	DeleteUser(email string) error
	Update(user User) error
}
