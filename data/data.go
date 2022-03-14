package data

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/alierkilic/do-cli/model"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func Open() error {
	var err error

	db, err := sql.Open("sqlite3", "./task.db")

	if err != nil {
		return err
	}

	Db = db
	return db.Ping()
}

func CreateTable() error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"task" TEXT NOT NULL,
		"done" INTEGER DEFAULT 0
	)`

	statement, err := Db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	statement.Exec()
	log.Println("Tasks table created")
	return nil
}

func SaveTask(task *model.NewTask) int64 {
	sqlStatement := `INSERT INTO tasks (task) VALUES($1) RETURNING id`
	id := 0
	err := Db.QueryRow(sqlStatement, &task.Task).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("New task added with ID: ", id)
	return int64(id)
}

func DeleteTask(id int) {
	sqlStatement := `DELETE FROM tasks WHERE id = $1`

	_, err := Db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Deleted task with ID: ", id)
}

func CompleteTask(id int) error {
	sqlStatement := `UPDATE tasks SET done=1 WHERE id = $1 and done=0 returning id`

	res := Db.QueryRow(sqlStatement, id)
	if res.Err() != nil {
		return res.Err()
	}
	var returnID int
	err := res.Scan(&returnID)
	if err != nil {
		fmt.Println("No task exists with the given ID...")
	}
	return err
}

func GetTasks() []model.Task {
	sqlStatement := `SELECT * FROM tasks WHERE done = 0`
	rows, err := Db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Task, &task.Done)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}

	log.Println("Got tasks")
	return tasks
}

func GetDoneTasks() []model.Task {
	sqlStatement := `SELECT * FROM tasks WHERE done=1`
	rows, err := Db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Task, &task.Done)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}

	log.Println("Got done tasks")
	return tasks
}
