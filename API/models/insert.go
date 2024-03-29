package models

import "github.com/gustavoddoki/MoneyTracker/API/db"

func AddTransaction(transaction Transaction) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	query := `INSERT INTO transactions (type, name, category, description, amount, date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err = conn.QueryRow(query, transaction.Type, transaction.Name, transaction.Category, transaction.Description, transaction.Amount, transaction.Date).Scan(&id)

	return
}
