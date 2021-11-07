package users

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        int32              `bson:"id"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Age       int32              `bson:"age"`
	Country   string             `bson:"country"`
	EntryDate primitive.DateTime `bson:"entryDate"`
}
