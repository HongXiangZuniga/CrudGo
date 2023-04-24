package users

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        int32              `bson:"id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Age       int32              `bson:"age" json:"age"`
	Country   string             `bson:"country" json:"country"`
	EntryDate primitive.DateTime `bson:"entryDate" json:"entryDate"`
}
