package main

import (
	"fmt"
	"html/template"
	"net/http"

	"example.com/todo-app/todo"
)

var serveUrl string = ":8090"

func main() {
	templ := template.Must(template.ParseFiles("./templ/layout.html", "./templ/tasks.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tasks := todo.GetTasks()
		if err := templ.Execute(w, tasks); err != nil {
			fmt.Println(err)
		}
	})
	fmt.Println("Listening on: ", serveUrl)
	http.ListenAndServe(serveUrl, nil)

}
