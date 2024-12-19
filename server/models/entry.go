package models

import (
    "go.mogodb.org/monogo-driver/bson/primitives"
)

type Entry struct {
	ID       primitives.ObjectID `bson:"_id,omitempty"`
	Dish *string `bson:"dish"`
	Fat *float64 `bson:"fat"`
	Ingedients *string `bson:"ingedients"`
	Calories *string `bson:"calories"`
}