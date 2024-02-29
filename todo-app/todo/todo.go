package todo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"iscompleted"`
}

func GetTasks() []Task {
	response, err := http.Get("http://localhost:8080/todo/tasks")

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	bodyBytes, _ := io.ReadAll(response.Body)

	var result []Task
	json.Unmarshal(bodyBytes, &result)
	return result
}
