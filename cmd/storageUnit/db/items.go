package db

import (
	"1michaelohayon/itemizer/typ"
	"fmt"
	"log"
)

type ItemsStore interface {
	InsertItem(typ.Item) error
	SelectOne(typ.Item, error)
	SelectAll() ([]typ.Item, error)
	UpdateOnesAmount(typ.Item) error
}

func CreateItemsTable() {
	createTable := `
        CREATE TABLE IF NOT EXISTS items(
            id INTEGER UNIQUE NOT NULL PRIMARY KEY,
            name STRING,
            amount INTEGER CHECK (amount >= 0)
    )`
	_, err := sqlDb.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("items ok")
	}
}

func InsertItem(item typ.Item) error {
	insert := `INSERT INTO items (id, name, amount) VALUES (?, ?, ?)`
	stmt, err := sqlDb.Prepare(insert)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(item.ID, item.Name, item.Amount)
	if err != nil {
		return err
	}
	return nil
}

func SelectAll() ([]typ.Item, error) {
	query := "SELECT * FROM items"
	rows, err := sqlDb.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	items := []typ.Item{}
	for rows.Next() {
		var item typ.Item
		err = rows.Scan(&item.ID, &item.Name, &item.Amount)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func SelectOne(id int64) (*typ.Item, error) {
	query := "SELECT * FROM items WHERE id = ?"
	stmt, err := sqlDb.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var item typ.Item
	err = stmt.QueryRow(id).Scan(&item.ID, &item.Name, &item.Amount)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func UpdateOnesAmount(item typ.Item) error {
	inDb, err := SelectOne(item.ID)
	if err != nil {
		return err
	}
	amount := inDb.Amount - item.Amount
	var syncErr error
	if amount < 0 {
		amount = 0
		syncErr = fmt.Errorf("sync error: less items are available in the db")
	}
	query := "UPDATE items SET amount = ? WHERE id = ?"
	_, err = sqlDb.Exec(query, amount, item.ID)
	if err != nil {
		return err
	}

	return syncErr
}

func DropTable() error {
	query := "DROP TABLE IF EXISTS items"
	_, err := sqlDb.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
