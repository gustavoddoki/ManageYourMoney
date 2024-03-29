package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gustavoddoki/ManageYourMoney/API/models"
)

func GetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	transactions, err := models.GetTransactions()
	if err != nil {
		log.Printf("Failed to get transactions.")
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}

func GetTransactionByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Failed to parse id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	transaction, err := models.GetTransactionByID(int64(id))
	if err != nil {
		log.Printf("Failed to get transaction: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}
