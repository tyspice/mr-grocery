package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Test struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Name  string             `bson:"name" json:"name"`
	Stoke string             `bson:"stoke" json:"stoke"`
}

type GroceryItem struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Item            string             `bson:"item" json:"item"`
	InventoryStatus int32              `bson:"inventoryStatus" json:"inventoryStatus"`
	Created         primitive.DateTime `bson:"created" json:"created"`
	Updated         primitive.DateTime `bson:"updated" json:"updated"`
	Notes           string             `bson:"notes" json:"notes"`
}

type NewGroceryItemRequest struct {
	Item            string `json:"item"`
	InventoryStatus int32  `json:"inventoryStatus"`
	Notes           string `json:"notes"`
}

type UpdateGroceryItem struct {
	Item            string             `bson:"item,omitempty" json:"item,omitempty"`
	InventoryStatus int32              `bson:"inventoryStatus,omitempty" json:"inventoryStatus,omitempty"`
	Notes           string             `bson:"notes,omitempty" json:"notes,omitempty"`
	Updated         primitive.DateTime `bson:"updated,omitempty" json:"updated,omitempty"`
}
