package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/breathingdust/house.api/db"
	"github.com/breathingdust/house.api/models"
	"github.com/gorilla/mux"
)

type (
	TransactionController struct {
		db *db.Database
	}
)

func NewTransactionController(db *db.Database) *TransactionController {
	return &TransactionController{db}
}

func (tc TransactionController) GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ts := tc.db.GetTransactions()

	j, _ := json.Marshal(ts)
	w.Write(j)
}

func (tc TransactionController) GetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	t := tc.db.GetTransaction(id)

	if t == nil {
		w.WriteHeader(404)
		return
	}

	uj, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(uj)
}

func (tc TransactionController) PostTransactionHandler(w http.ResponseWriter, r *http.Request) {
	t := models.Transaction{}
	json.NewDecoder(r.Body).Decode(&t)

	ct := tc.db.CreateTransaction(&t)

	tj, _ := json.Marshal(ct)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(tj)
}
