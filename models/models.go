package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Test struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Name  string             `bson:"name" json:"name"`
	Stoke string             `bson:"stoke" json:"stoke"`
}
