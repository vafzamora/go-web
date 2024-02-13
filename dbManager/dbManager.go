package dbManager

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/mysql?parseTime=true")

	logError(err)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func CloseConnection(db *sql.DB) {
	db.Close()
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InitDb() {
	db := OpenConnection()
	defer CloseConnection(db)

	var err error
	createDbCommand := "CREATE DATABASE IF NOT EXISTS todo"
	_, err = db.Exec(createDbCommand)
	logError(err)

	createTableCommand := "CREATE TABLE IF NOT EXISTS todo.tasks (`id` INT AUTO_INCREMENT, `title` VARCHAR(50) NOT NULL, `iscompleted` BOOL, PRIMARY KEY (ID));"
	_, err = db.Exec(createTableCommand)
	logError(err)

	var hasTasks bool

	if err := db.QueryRow("SELECT EXISTS(SELECT * FROM todo.tasks)").Scan(&hasTasks); err == nil && !hasTasks {
		createTasksCommand := "INSERT todo.tasks (title, iscompleted) VALUES (?,?)"
		for i := 1; i <= 5; i++ {
			_, err := db.Exec(createTasksCommand, fmt.Sprintf("Task %v", i), 0)
			logError(err)
		}
	}
}
