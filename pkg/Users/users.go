package users

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        int                `bson:"id"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Age       int                `bson:"age"`
	Country   string             `bson:"country"`
	EntryDate primitive.DateTime `bson:"entryDate"`
}
