package models

import "github.com/gustavoddoki/MoneyTracker/API/db"

func GetTransactionByID(id int64) (transaction Transaction, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	query := `SELECT * FROM transactions WHERE id=$1`
	row := conn.QueryRow(query, id)

	err = row.Scan(&transaction.ID, &transaction.Type, &transaction.Name, &transaction.Category, &transaction.Description, &transaction.Amount, &transaction.Date, &transaction.RegistryTime)

	return
}

func GetTransactions() (transactions []Transaction, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	query := `SELECT * FROM transactions`
	rows, err := conn.Query(query)
	if err != nil {
		return
	}

	for rows.Next() {
		var transaction Transaction

		err = rows.Scan(&transaction.ID, &transaction.Type, &transaction.Name, &transaction.Category, &transaction.Description, &transaction.Amount, &transaction.Date, &transaction.RegistryTime)
		if err != nil {
			continue
		}

		transactions = append(transactions, transaction)

	}
	return
}
