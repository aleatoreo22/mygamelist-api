package database

import (
	"database/sql"
)

var strConnection string
var dataBase *sql.DB

func Open() error {
	if dataBase.Stats().OpenConnections > 0 {
		Close()
	}
	db, err := sql.Open("mysql", strConnection)
	dataBase = db
	return err
}

func Close() {
	dataBase.Close()
}

func Select(query string) (*sql.Rows, error) {
	rows, err := dataBase.Query(query)
	if err != nil {
		return rows, err
	}
	defer rows.Close()
	return rows, err
}

func Execute(query string) error {
	_, err := dataBase.Exec(query)
	return err
}
