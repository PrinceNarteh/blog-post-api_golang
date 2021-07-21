package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Title   string             `json:"title"`
	Content string             `json:"content"`
}
