package main

import (
	"log"
	"net/http"
	"os"

	"github.com/breathingdust/houseapp/controllers"
	"github.com/breathingdust/houseapp/db"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)

	var mgoConn = os.Getenv("MGOCONN")

	if mgoConn == "" {
		log.Fatal("No connection string found.")
	}

	log.Println(mgoConn)

	db := db.NewDatabase(mgoConn, "house")

	tc := controllers.NewTransactionController(db)
	hc := controllers.NewHomeController()

	r.HandleFunc("/", hc.HomeHandler).Methods("GET")
	r.HandleFunc("/transactions", tc.GetTransactionsHandler).Methods("GET")
	r.HandleFunc("/transactions/{id}", tc.GetTransactionHandler).Methods("GET")
	r.HandleFunc("/transactions", tc.PostTransactionHandler).Methods("POST")
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
