package tests

import (
	"errors"
	"net/http"
	"prac-go-mvc-rest/models/entities"
)

type MockTodoController struct {
}

func (mtc *MockTodoController) GetTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (mtc *MockTodoController) PostTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
}

func (mtc *MockTodoController) PutTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}

func (mtc *MockTodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}

type MockTodoRepository struct {
}

func (mtr *MockTodoRepository) GetTodo() (todos []entities.TodoEntity, err error) {
	todos = []entities.TodoEntity{}
	return
}

func (mtr *MockTodoRepository) InsertTodo(todo entities.TodoEntity) (id int, err error) {
	id = 2
	return
}

type MockTodoRepositoryGetExist struct {
}

func (mtrgex *MockTodoRepositoryGetExist) GetTodo() (todos []entities.TodoEntity, err error) {
	todos = []entities.TodoEntity{}
	todos = append(todos, entities.TodoEntity{Id: 1, Title: "title1", Content: "contents1"})
	todos = append(todos, entities.TodoEntity{Id: 2, Title: "title2", Content: "contents2"})
	return
}

func (mtrgex *MockTodoRepositoryGetExist) InsertTodo(todos entities.TodoEntity) (id int, err error) {
	return
}

type MockTodoRepositoryError struct {
}

func (mtrge *MockTodoRepositoryError) GetTodo() (todos []entities.TodoEntity, err error) {
	err = errors.New("unexpected error occurred")
	return
}

func (mtrie *MockTodoRepositoryError) InsertTodo(todo entities.TodoEntity) (id int, err error) {
	err = errors.New("unexpected error occurred")
	return
}
