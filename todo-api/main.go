package main

import (
	"fmt"
	"net/http"
	"strings"

	"example.com/todo-api/dbManager"
	"example.com/todo-api/todo"
	"github.com/gorilla/mux"
)

func main() {
	dbManager.InitDb()

	router := mux.NewRouter()
	todo.MapTodoHandlers(router)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to my website!")
	})

	fmt.Println("Start Listen")
	if err := http.ListenAndServe(":8080", trailingSlashRemovalMiddleware(router)); err != nil {
		fmt.Println(err)
	}
}

func trailingSlashRemovalMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}
