package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDB() error {
	log.Println("connectDB")

	if DB != nil {
		return nil
	}

	var err error
	DB, err = sql.Open("sqlite3", "./mydb.db")
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	log.Println("DB connected")

	return nil
}
