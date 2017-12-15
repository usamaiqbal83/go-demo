package model

import "gopkg.in/mgo.v2/bson"

type (
	UserRole struct {
		ID          bson.ObjectId `json:"id" bson:"_id"`
		Name        string        `json:"name" bson:"name"`
		Description string        `json:"desc" bson:"desc"`
	}
)

type UserRoles []UserRole
