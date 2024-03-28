package db

import (
	"database/sql"
	"fmt"

	"github.com/gustavoddoki/ManageYourMoney/API/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	query := fmt.Sprintf("host-%s port-%s user-%s password-%s dbname-%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", query)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}
