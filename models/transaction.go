package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	Transaction struct {
		Id        bson.ObjectId `json:"id" bson:"_id"`
		Name      string        `json:"name" bson:"name"`
		Timestamp time.Time     `json:"Timestamp" bson:"timestamp"`
	}
)

type Transactions []Transaction
