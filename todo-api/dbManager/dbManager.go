package dbManager

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const filename = "todo.db"

func OpenConnection() *sql.DB {
	db, err := sql.Open("sqlite3", filename)

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

	createTableCommand := "CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, title VARCHAR(50) NOT NULL, iscompleted BOOL);"
	_, err = db.Exec(createTableCommand)
	logError(err)

	var hasTasks bool

	if err := db.QueryRow("SELECT EXISTS(SELECT * FROM tasks)").Scan(&hasTasks); err == nil && !hasTasks {
		createTasksCommand := "INSERT INTO tasks (title, iscompleted) VALUES (?,?)"
		for i := 1; i <= 5; i++ {
			_, err := db.Exec(createTasksCommand, fmt.Sprintf("Task %v", i), 0)
			logError(err)
		}
	}
}
