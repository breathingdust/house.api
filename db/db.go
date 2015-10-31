package db

import (
	"log"

	"github.com/breathingdust/houseapp/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Database struct {
	session      *mgo.Session
	databaseName string
}

func NewDatabase(conn string, databaseName string) *Database {
	db := Database{nil, databaseName}
	db.init(conn)
	return &db
}

func (database *Database) init(conn string) {
	s, err := mgo.Dial(conn)

	if err != nil {
		panic(err)
	}

	database.session = s
}

func (database *Database) getSession() *mgo.Session {
	return database.session.Copy()
}

func (database *Database) GetTransactions() *models.Transactions {
	ts := models.Transactions{}
	database.getSession().DB(database.databaseName).C("transactions").Find(nil).All(&ts)
	return &ts
}

func (database *Database) GetTransaction(id string) *models.Transaction {
	oid := bson.ObjectIdHex(id)
	t := models.Transaction{}
	database.getSession().DB(database.databaseName).C("transactions").FindId(oid).One(&t)
	return &t
}

func (database *Database) CreateTransaction(t *models.Transaction) *models.Transaction {
	t.Id = bson.NewObjectId()

	session := database.getSession()

	log.Println(t.Name)

	session.DB(database.databaseName).C("transactions").Insert(t)

	ct := models.Transaction{}
	session.DB(database.databaseName).C("transactions").FindId(t.Id).One(&ct)

	return &ct
}
