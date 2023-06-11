package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	user   = "postgres"
	dbName = "postgres"
)

var (
	psqlDb *sql.DB
	pass   string
	host   = "localhost"
	dbPort = "6432"
)

func init() {
	if pass = os.Getenv("PSQL_PASS"); len(pass) == 0 {
		log.Fatal("missing psql password. env: PSQL_PASS")
	}

	hostEnv := os.Getenv("PSQL_HOST")
	portEnv := os.Getenv("PSQL_PORT")
	if len(hostEnv) > 0 {
		host = hostEnv
	}
	if len(portEnv) > 0 {
		dbPort = portEnv
	}

}

func ConnectToPSQL() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, dbPort, user, pass, dbName)
	fmt.Printf("Connecting to the database at: %s:%s as %s\n", host, dbPort, user)
	db, err := sql.Open("postgres", connStr)
	psqlDb = db
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return nil
	}
	if err := ping(psqlDb); err != nil {
		fmt.Println("DB Error:", err.Error())
	} else {
		fmt.Println("db connection ok")
	}
	return db
}

func ping(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return err
	}
	return nil
}
