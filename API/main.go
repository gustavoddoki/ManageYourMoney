package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// Transaction represents a financial transaction, either an expense or an income
type Transaction struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	Name         string `json:"name"`
	Category     string `json:"category"`
	Description  string `json:"description"`
	Amount       string `json:"amount"`
	Date         string `json:"date"`
	RegistryTime string `json:"registryTime"`
}

// Transaction is a slice of Transsaction
type Transactions []Transaction

// Create a new database if not exists
func CreateDatabase() {
	db, err := sql.Open("sqlite3", "./finances.db")

	if err != nil {
		log.Fatal(err)
	}

	query := "CREATE TABLE IF NOT EXISTS transactions (id INTEGER PRIMARY KEY, type VARCHAR, name VARCHAR, description TEXT, amount REAL, category VARCHAR, date DATE, registry_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP)"
	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err)
	}
}

func asfasdfsadf(c *gin.Context) {

	db, err := sql.Open("sqlite3", "./finances.db")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open the database"})
		return
	}
	defer db.Close()

	query := "SELECT * FROM transactions"
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute the query"})
		return
	}
	defer rows.Close()

	transactions := []Transaction{}
	for rows.Next() {
		transaction := Transaction{}
		err = rows.Scan(&transaction.ID, &transaction.Date, &transaction.Description, &transaction.Amount, &transaction.Category)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan the row"})
			return
		}

		transactions = append(transactions, transaction)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve the rows"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func main() {

	os.Setenv("CGO_ENABLED", "1")

	CreateDatabase()

	r := gin.Default()
	r.GET("/transactions", getTransactions)
	r.Run("localhost:3000")

}
