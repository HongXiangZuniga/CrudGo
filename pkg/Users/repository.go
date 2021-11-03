package users

type UsersMongoRepo interface {
	FindUser(id int) (*User, error)
	CreateUser(user User) error
	DeleteUser(email string) error
	Update(user User) error
}
