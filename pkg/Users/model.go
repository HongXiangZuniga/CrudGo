package users

type User struct {
	Name  string `bson:"Name"`
	Email string `bson:"Email"`
	Age   int    `bson:"Age"`
}
