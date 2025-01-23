package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"server/src"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDB() {
	file, err := os.Create("sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()

	db, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	createTable := `
		create table if not exists quotation (
			id integer not null primary key autoincrement,
			bid real,
			create_at text
		);
	`

	statement, err := db.Prepare(createTable)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
}

func InsertQuotation(quotation *src.Quotation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	db, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return err
	}
	defer db.Close()

	insertQuotation := `insert into quotation (bid, create_at) values(?, ?)`

	statement, err := db.PrepareContext(ctx, insertQuotation)
	if err != nil {
		return err
	}

	bid, err := strconv.ParseFloat(quotation.Bid, 64)
	if err != nil {
		return err
	}

	statement.Exec(bid, time.Now().String())

	return nil
}
