package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RecipeModel struct
type RecipeModel struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name         string             `bson:"name" json:"name,omitempty"`
	ImageUrl     string             `bson:"imageUrl" json:"imageUrl,omitempty"`
	Category     string             `bson:"category" json:"category,omitempty"`
	Description  string             `bson:"description" json:"description,omitempty"`
	Instructions string             `bson:"instructions" json:"instructions,omitempty"`
	CreatedDate  time.Time          `bson:"createdDate,omitempty" json:"createdDate,omitempty"`
	Likes        int                `bson:"likes,omitempty" json:"likes,omitempty"`
	Username     string             `bson:"username,omitempty" json:"username,omitempty"`
}

// UserModel struct
type UserModel struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Username  string             `bson:"username" json:"username,omitempty"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Email     string             `bson:"email" json:"email,omitempty"`
	JoinDate  time.Time          `bson:"joinDate" json:"joinDate,omitempty"`
	Favorites []interface{}      `bson:"favorites" json:"favorites,omitempty" model:"RecipeModel"`
	Token     string             `json:"token,omitempty"` // graphql only
}
