package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/pustaka-api")
	if err != nil {
		return nil, err
	}

	return db, nil
}
