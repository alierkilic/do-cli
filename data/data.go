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
	dropTableSQL := `DROP TABLE IF EXISTS tasks`
	_, err := Db.Exec(dropTableSQL)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"task" TEXT NOT NULL,
		"done" BOOLEAN DEFAULT FALSE,
		"daily" BOOLEAN DEFAULT FALSE
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
	sqlStatement := `INSERT INTO tasks (task, daily) VALUES($1, $2) RETURNING id`
	id := 0
	err := Db.QueryRow(sqlStatement, &task.Task, &task.Daily).Scan(&id)
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
	sqlStatement := `UPDATE tasks SET done=TRUE WHERE id = $1 and done=FALSE returning id`

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

func GetTasks(daily bool) []model.Task {
	sqlStatement := `SELECT * FROM tasks WHERE done = FALSE AND daily = $1`
	rows, err := Db.Query(sqlStatement, daily)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Task, &task.Done, &task.Daily)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}

	log.Println("Got tasks")
	return tasks
}

func GetDoneTasks() []model.Task {
	sqlStatement := `SELECT * FROM tasks WHERE done=TRUE`
	rows, err := Db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Task, &task.Done, &task.Daily)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}

	log.Println("Got done tasks")
	return tasks
}
