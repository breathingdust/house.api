package db

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/breathingdust/house.api/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
)

func clearDatabase(conn string) {
	s, _ := mgo.Dial(conn)
	//defer s.Close()
	s.DB("house_test").C("transactions").RemoveAll(nil)
}

func getConnectionString() string {
	var mgoConn = os.Getenv("MGOCONN")
	if mgoConn == "" {
		log.Fatal("No connection string found.")
	}
	return mgoConn
}

func TestGetTransaction(t *testing.T) {
	mgoConn := getConnectionString()
	clearDatabase(mgoConn)
	database := NewDatabase(mgoConn, "house_test")

	tr := models.Transaction{}

	tr.Name = "Transaction Name"
	tr.Timestamp = time.Now()
	tr.Amount = 1.50
	tr.Buyer = "simon"
	tr.Type = "shared"

	createdTransaction := database.CreateTransaction(&tr)

	gotTransaction := database.GetTransaction(createdTransaction.Id.Hex())

	assert.Equal(t, tr.Name, gotTransaction.Name)
}

func TestGetTransactions(t *testing.T) {
	mgoConn := getConnectionString()
	clearDatabase(mgoConn)
	database := NewDatabase(mgoConn, "house_test")

	tr := models.Transaction{}

	tr.Name = "Transaction Name"
	tr.Timestamp = time.Now()
	tr.Amount = 1.50
	tr.Buyer = "simon"
	tr.Type = "shared"

	database.CreateTransaction(&tr)

	transactions := database.GetTransactions()

	assert.Equal(t, tr.Name, (*transactions)[0].Name)
}
