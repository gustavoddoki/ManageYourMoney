package models

import "github.com/gustavoddoki/ManageYourMoney/API/db"

func Update(id int64, transaction Transaction) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	query := "UPDATE transactions SET type=$2, name=$3, category=$4, description=$5, amount=$6, date=$7 WHERE id=$1"
	res, err := conn.Exec(query, id, transaction.Type, transaction.Name, transaction.Category, transaction.Description, transaction.Amount, transaction.Date)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
