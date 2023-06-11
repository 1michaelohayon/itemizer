package db

import (
	"1michaelohayon/itemizer/typ"
	"fmt"
)

func CreateItemDataTable() {
	createTable := `
        CREATE TABLE IF NOT EXISTS items_data(
            id BIGSERIAL PRIMARY KEY,
            item_id BIGINT NOT NULL,
            name VARCHAR(50),
            amount INT,
            errors INT,
            sender_id BIGINT NOT NULL,
            sender_timestamp TIMESTAMP,
            storage_unit_id VARCHAR(50)
    )`

	_, err := psqlDb.Exec(createTable)
	if err != nil {
		fmt.Println("Error creating database:", err)
	} else {
		fmt.Println("items_data ok")
	}
}

func InsertItem(data typ.ItemData) error {
	insert := `
    INSERT INTO items_data (
    item_id,
    name,
    amount,
    errors,
    sender_id,
    sender_timestamp,
    storage_unit_id) VALUES($1, $2, $3, $4, $5, $6, $7)`
	stmt, err := psqlDb.Prepare(insert)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		data.Item.ID,
		data.Item.Name,
		data.Item.Amount,
		len(data.Item.Errors),
		data.Item.Sender.ID,
		data.Item.Sender.Time,
		data.StorageUnit.ID,
	)
	if err != nil {
		return err
	}
	fmt.Printf("Inserted: %v\n", data)
	return nil
}

func SelectAll() ([]typ.ItemData, error) {
	query := "SELECT item_id, name, amount, sender_id, storage_unit_id FROM items_data"
	rows, err := psqlDb.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	itemsData := []typ.ItemData{}
	for rows.Next() {
		var (
			itemD typ.ItemData
		)
		err = rows.Scan(&itemD.Item.ID, &itemD.Item.Name, &itemD.Item.Amount,
			&itemD.Item.Sender.ID, &itemD.StorageUnit.ID)

		if err != nil {
			return nil, err
		}
		itemsData = append(itemsData, itemD)
	}
	return itemsData, nil
}
