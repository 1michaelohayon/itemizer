package db

import (
	"1michaelohayon/itemizer/typ"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	err := Connect("./TEST_data.db")
	if err != nil {
		log.Fatal("connection to the database fail")
	}
	if err := DropTable(); err != nil {
		log.Fatal(err)
	}
	CreateItemsTable()
}

func tearDown() {
	query := "DELETE FROM items"
	if _, err := sqlDb.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func TestInsert(t *testing.T) {
	defer tearDown()
	dummy := typ.Item{ID: 4, Name: "test1", Amount: 1}
	if err := InsertItem(dummy); err != nil {
		t.Fatal("failed inserting new item", dummy)
	}

	items, err := SelectAll()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(items))
	assert.Equal(t, items[0], dummy)
}

func TestSelectAll(t *testing.T) {
	defer tearDown()
	dummys := []typ.Item{
		{ID: 1, Name: "test1", Amount: 1},
		{ID: 2, Name: "test2", Amount: 6},
		{ID: 3, Name: "test3", Amount: 2},
		{ID: 4, Name: "test4", Amount: 10},
	}
	for _, d := range dummys {
		if err := InsertItem(d); err != nil {
			t.Fatal("failed inserting new item", d)
		}
	}

	items, err := SelectAll()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, len(dummys), len(items))
	assert.Equal(t, dummys, items)
}

func TestSelectOne(t *testing.T) {
	defer tearDown()

	dummy := typ.Item{ID: 4, Name: "test1", Amount: 1}
	notDummy := typ.Item{ID: 5, Name: "test7", Amount: 3}

	if err := InsertItem(dummy); err != nil {
		t.Fatal("failed inserting new item", dummy)
	}

	item, err := SelectOne(dummy.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, *item, dummy)
	assert.NotEqual(t, *item, notDummy)
}

func TestUpdateAmount(t *testing.T) {
	defer tearDown()

	dummy := typ.Item{ID: 4, Name: "test1", Amount: 3}

	if err := InsertItem(dummy); err != nil {
		t.Fatal("failed inserting new item", dummy)
	}

	updated := dummy
	updated.Amount = 2
	if err := UpdateOnesAmount(updated); err != nil {
		t.Fatal(err)
	}

	fromDb, err := SelectOne(dummy.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, fromDb, dummy)
	assert.Equal(t, fromDb.Amount, dummy.Amount-2)

	updated.Amount = 51
	err = UpdateOnesAmount(updated)
	assert.Equal(t, err.Error(), "sync error: less items are available in the db")

	fromDb, err = SelectOne(dummy.ID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 0, fromDb.Amount)

}
