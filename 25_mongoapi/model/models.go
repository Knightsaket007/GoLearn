package models;

import(
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Netflix struct{
	ID bson.ObjectID `json: "_id, omitempty" beson:"_id, omitempty"`
}