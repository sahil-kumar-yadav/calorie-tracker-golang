package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// go lang do not understand goland
type Entry struct {
	ID          primitive.ObjectID `bson:"id"`
	Dish        *string            `json:"dish"`
	Fat         *float64           `json:"fat"`
	Ingredients *string            `json:"ingredients"`
	Calories    *string            `json:"calories"`
}