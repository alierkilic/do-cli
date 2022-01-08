package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/sqlite3"
)

var db *sql.DB

func Open() error {
	var err error

	db, err := sql.Open("sqlite3", "./task.db")

	if err != nil {
		return err
	}

	return db.Ping()
}

func CreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"task" TEXT,
		"done" BOOLEAN
	)`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Tasks table created")
}
