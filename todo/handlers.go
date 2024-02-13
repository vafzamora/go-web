package todo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func MapTodoHandlers(router *mux.Router) {
	router.HandleFunc("/todo/tasks", func(w http.ResponseWriter, r *http.Request) {
		tasks := GetTasks()
		json.NewEncoder(w).Encode(tasks)
	}).Methods("GET")

	router.HandleFunc("/todo/tasks/{id:\\d+}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		if res := GetTask(id); res.Id != id {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Task %d not found.", id)
		} else {
			json.NewEncoder(w).Encode(res)
		}
	}).Methods("GET")

	router.HandleFunc("/todo/tasks", func(w http.ResponseWriter, r *http.Request) {
		var newTask Task
		json.NewDecoder(r.Body).Decode(&newTask)

		CreateTask(&newTask)
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(newTask)
	}).Methods("POST")

	router.HandleFunc("/todo/tasks/{id:\\d+}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])

		var task Task
		json.NewDecoder(r.Body).Decode(&task)

		if task.Id != id {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Incorrect task.")
		}

		UpdateTask(&task)
		//w.WriteHeader(http.StatusNoContent)

	}).Methods("PUT")

	router.HandleFunc("/todo/tasks/{id:\\d+}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["id"])

		DeleteTask(id)
		//w.WriteHeader(http.StatusNoContent)

	}).Methods("DELETE")
}
