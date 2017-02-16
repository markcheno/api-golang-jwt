package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Permission model
type Permission struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	EndPoint  string        `json:"endpoint" bson:"endpoint"`
	Perms     []PermBitwise `json:"perms" bson:"perms"`
	Owner     bson.ObjectId `json:"owner" bson:"owner"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}

//PermBitwise methods request and values
type PermBitwise struct {
	Method string `json:"method" bson:"method"`
	Value  int    `json:"value" bson:"value"`
}
