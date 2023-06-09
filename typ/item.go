package typ

import "time"

type Item struct {
	ID     int64
	Name   string
	Amount int
	Errors []string
	Sender Sender
}

type Sender struct {
	ID   int64
	Time time.Time
}

type StorageUnit struct {
	ID string
}

type ItemData struct {
	Item        Item
	StorageUnit StorageUnit
}
