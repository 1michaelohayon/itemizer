package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var sqlDb *sql.DB

func Connect(path string) error {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return err
	} else {
		fmt.Println("db ok")
	}
	sqlDb = db
	return nil
}
