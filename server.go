package main

import (
	"log"
	"net/http"
	"os"

	"github.com/breathingdust/house.api/controllers"
	"github.com/breathingdust/house.api/db"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	log.Println("Listening on port 8080")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
