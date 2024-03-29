package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func OpenConnection() (*sql.DB, error) {

	conn, err := sql.Open("sqlite3", "./finances.db")
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}

func CreateDatabase() {

	query := "CREATE TABLE IF NOT EXISTS transactions (id INTEGER PRIMARY KEY, type VARCHAR, name VARCHAR, category VARCHAR, description TEXT, amount REAL, date DATE, time_registry TIMESTAMP DEFAULT CURRENT_TIMESTAMP)"
	conn, err := OpenConnection()

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = conn.Exec(query)
	if err != nil {
		panic(err)
	}

}
