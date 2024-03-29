package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gustavoddoki/MoneyTracker/API/models"
)

func DeleteTransactionHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.Replace(r.URL.Path, "/", "", 1))
	if err != nil {
		log.Printf("Failed to parse id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.DeleteTransanction(int64(id))
	if err != nil {
		log.Printf("Failed to delete transaction: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("WARNING: %v transactions was deleted.", rows)
	}

	res := map[string]any{
		"Error":   false,
		"Message": "Transaction deleted successfully!",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
