package main

import (
	"fmt"
	"net/http"
	"strings"

	"example.com/go-web/dbManager"
	"example.com/go-web/todo"
	"github.com/gorilla/mux"
)

func main() {

	dbManager.InitDb()

	router := mux.NewRouter()
	//router.Use(trailingSlashRemovalMiddleware)
	todo.MapTodoHandlers(router)

	// r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
	// 	vars := mux.Vars(r)
	// 	title := vars["title"]
	// 	page := vars["page"]
	// 	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	// })

	fs := http.FileServer(http.Dir("static/"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

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
