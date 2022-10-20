package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"fmt"
	"prac-go-mvc-rest/controllers/dto"
	"prac-go-mvc-rest/models/entities"
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
	todos, err := tc.tr.GetTodo()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var todoResponses []dto.TodoResponse
	for _, v := range todos {
		todoResponses = append(todoResponses, dto.TodoResponse{Id: v.Id, Title: v.Title, Content: v.Content})
	}

	var todoResponse dto.TodosResponse
	todoResponse.Todos = todoResponses

	output, _ := json.MarshalIndent(todoResponse.Todos, "", "\t\t")
	w.Header().Set("Content-type", "application/json")
	w.Write(output)
}

func (tc *todoController) PostTodo(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var todoRequest dto.TodoRequest
	json.Unmarshal(body, &todoRequest)

	todo := entities.TodoEntity{Title: todoRequest.Title, Content: todoRequest.Content}

	id, err := tc.tr.InsertTodo(todo)
	if err != nil {
		log.Print(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(id))
	w.WriteHeader(201)
}

func (tc *todoController) PutTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("put")
}

func (tc *todoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete")
}
