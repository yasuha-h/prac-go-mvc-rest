package controllers

import (
	"net/http"

	"fmt"
	"prac-go-mvc-rest/models/repositories"
)

type TodoController interface {
	GetTodo(w http.ResponseWriter, r *http.Request)
	PostTodo(w http.ResponseWriter, r *http.Request)
	PutTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}

type todoController struct {
	tr repositories.TodoRepository
}

func NewTodoController(tr repositories.TodoRepository) TodoController {
	return &todoController{tr}
}

func (tc *todoController) GetTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get")
}

func (tc *todoController) PostTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post")
}

func (tc *todoController) PutTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("put")
}

func (tc *todoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete")
}
