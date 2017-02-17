package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Project model
type Project struct {
	ID          bson.ObjectId   `json:"id" bson:"_id"`
	Label       string          `json:"label" bson:"label"`
	Slug        string          `json:"slug" bson:"slug"`
	Description string          `json:"description" bson:"description"`
	Owner       bson.ObjectId   `json:"owner" bson:"owner"`
	Users       []bson.ObjectId `json:"users" bson:"users"`
	CreatedAt   time.Time       `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at" bson:"updated_at"`
}
