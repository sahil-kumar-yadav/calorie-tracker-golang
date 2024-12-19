package models

import (
	"go.mogodb.org/monogo-driver/bson/primitives"
)

// go lang do not understand goland
type Entry struct {
	ID         primitives.ObjectID `bson:"_id,omitempty"`
	Dish       *string             `json:"dish"`
	Fat        *float64            `json:"fat"`
	Ingedients *string             `json:"ingedients"`
	Calories   *string             `json:"calories"`
}
