package model

import (
	"go.mongodb.org/mongo-driver/v2/bson/primitive"
)

type Netflix struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Movie   string             `bson:"movie"`
	Watched bool               `bson:"watched"`
}