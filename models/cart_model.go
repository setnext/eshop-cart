package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cart struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	UserId    string             `json:"userId,omitempty" validate:"required"`
	CartItems *[]CartItem        `json:"cartItems,omitempty" validate:"required"`
}

type CartItem struct {
	ItemNo      int    `json:"itemNo,omitempty" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required"`
	Name        string `json:"productName" validate:"required"`
	Description string `json:"productDescription" validate:"required"`
	Category    string `json:"category" validate:"required"`
	ImageURL    string `json:"imageUrl" validate:"required"`
	ProductUrl  string `json:"productUrl" validate:"required"`
}
