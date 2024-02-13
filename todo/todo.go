package todo

import (
	"log"

	"example.com/go-web/dbManager"
)

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"iscompleted"`
}

func GetTasks() []Task {
	db := dbManager.OpenConnection()
	defer dbManager.CloseConnection(db)

	rows, err := db.Query("SELECT id, title, iscompleted FROM todo.tasks")
	logError(err)
	defer rows.Close()

	var result []Task
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.Id, &t.Title, &t.IsCompleted)
		logError(err)
		result = append(result, t)
	}
	logError(rows.Err())
	return result
}

func GetTask(id int) Task {
	db := dbManager.OpenConnection()
	defer dbManager.CloseConnection(db)

	var t Task
	err := db.QueryRow("SELECT id, title, iscompleted FROM todo.tasks WHERE id=?", id).Scan(&t.Id, &t.Title, &t.IsCompleted)
	logError(err)

	return t
}

func CreateTask(newTask *Task) {
	db := dbManager.OpenConnection()
	defer dbManager.CloseConnection(db)

	result, err := db.Exec("INSERT todo.tasks (title, iscompleted) VALUES (?,?)", newTask.Title, newTask.IsCompleted)
	logError(err)

	id, err := result.LastInsertId()
	newTask.Id = int(id)

	logError(err)
}

func UpdateTask(task *Task) {
	db := dbManager.OpenConnection()
	defer dbManager.CloseConnection(db)

	_, err := db.Exec("UPDATE todo.tasks SET title = ?, iscompleted = ? WHERE id = ?", task.Title, task.IsCompleted, task.Id)
	logError(err)
}

func DeleteTask(id int) {
	db := dbManager.OpenConnection()
	defer dbManager.CloseConnection(db)

	_, err := db.Exec("DELETE FROM todo.tasks WHERE id = ?", id)
	logError(err)
}

func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}
