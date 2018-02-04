package models

import "gopkg.in/mgo.v2/bson"

// the properties in mongodb document
type Todo struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Content        string        `bson:"content" json:"content"`
}
