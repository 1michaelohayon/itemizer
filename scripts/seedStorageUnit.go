package main

import (
	"1michaelohayon/itemizer/cmd/storageUnit/db"
	"1michaelohayon/itemizer/typ"
)

func main() {
	db.Connect("./data.db")
	db.DropTable()
	db.CreateItemsTable()

	items := []typ.Item{
		{ID: 1, Name: "Paper 12 x 9", Amount: 73},
		{ID: 2, Name: "Paper 8 x 6", Amount: 73},
		{ID: 3, Name: "Paper", Amount: 73},
		{ID: 4, Name: "Box 2 x 5", Amount: 200},
		{ID: 5, Name: "Box 4 x 4", Amount: 170},
		{ID: 6, Name: "Box 9 x 10", Amount: 170},
		{ID: 7, Name: "Stapler", Amount: 83},
		{ID: 8, Name: "Towels", Amount: 143},
		{ID: 9, Name: "Sponge", Amount: 100},
	}

	for _, item := range items {
		db.InsertItem(item)
	}
}
