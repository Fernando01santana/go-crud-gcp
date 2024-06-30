package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Quantity int                `json:"quantity,omitempty" bson:"quantity,omitempty"`
}
