package todo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"iscompleted"`
}

var apiBaseAddress string = "http://localhost:8080"

func init() {
	if tmp := os.Getenv("TODOAPI_BASEADDRESS"); tmp != "" {
		apiBaseAddress = strings.TrimSuffix(tmp, "/")
	}
	fmt.Printf("API base address: %s\n", apiBaseAddress)
}

func GetTasks() []Task {
	response, err := http.Get(fmt.Sprintf("%s/todo/tasks", apiBaseAddress))

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	bodyBytes, _ := io.ReadAll(response.Body)

	var result []Task
	json.Unmarshal(bodyBytes, &result)
	return result
}
