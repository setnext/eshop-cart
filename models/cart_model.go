package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cart struct {
	Id        primitive.ObjectID `bson:"id,omitempty"`
	UserId    string             `json:"userId" bson:"userId" validate:"required"`
	CartItems []CartItem         `json:"cartItems" bson:"cartItems" validate:"required"`
}

type CartItem struct {
	ItemNo      int    `json:"itemNumber" bson:"itemNumber" validate:"required"`
	Quantity    int    `json:"quantity" bson:"quantity" validate:"required"`
	Name        string `json:"productName" bson:"productName" validate:"required"`
	Description string `json:"productDescription" bson:"productDescription" validate:"required"`
	Category    string `json:"category" bson:"category" validate:"required"`
	ImageURL    string `json:"imageUrl" bson:"imageUrl" validate:"required"`
	ProductUrl  string `json:"productUrl" bson:"productUrl" validate:"required"`
	UnitPrice   int    `json:"unitPrice" bson:"unitPrice" validate:"required"`
	TotalPrice  int    `json:"totalPrice" bson:"totalPrice" validate:"required"`
}
