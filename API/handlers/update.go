package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gustavoddoki/MoneyTracker/API/models"
)

func UpdateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.Replace(r.URL.Path, "/", "", 1))
	if err != nil {
		log.Printf("Failed to parse id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var transaction models.Transaction

	err = json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		log.Printf("Failed to decode json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.UpdateTransaction(int64(id), transaction)
	if err != nil {
		log.Printf("Failed to update transaction: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("WARNING: %v transactions was updated.", rows)
	}

	res := map[string]any{
		"Error":   false,
		"Message": "Transaction updated successfully!",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
