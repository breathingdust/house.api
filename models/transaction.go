package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	Transaction struct {
		Id        bson.ObjectId `json:"id" bson:"_id"`
		Name      string        `json:"name" bson:"name"`
		Amount    float64       `json:"Amount" bson:"amount"`
		Buyer     string        `json:"Buyer" bson:"buyer"`
		Type      string        `json:"Type" bson:"type"`
		Timestamp time.Time     `json:"Timestamp" bson:"timestamp"`
	}
)

type Transactions []Transaction
