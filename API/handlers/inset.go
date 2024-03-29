package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gustavoddoki/ManageYourMoney/API/models"
)

func HandlerAddTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction

	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		log.Printf("Failed to decode json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.AddTransaction(transaction)

	var res map[string]any

	if err != nil {
		res = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Failed to add transaction: %v", err),
		}
	} else {
		res = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Transaction added successfully! ID %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
