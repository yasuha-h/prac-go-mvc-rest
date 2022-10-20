package main

import (
	"net/http"

	"prac-go-mvc-rest/controllers"
	"prac-go-mvc-rest/models/repositories"
)

var tr = repositories.NewTodoRepository()
var tc = controllers.NewTodoController(tr)
var ro = controllers.NewRouter(tc)

func main() {
	server := http.Server{
		Addr: ":8000",
	}
	http.HandleFunc("/", ro.HandleTodosRequest)
	server.ListenAndServe()
}
