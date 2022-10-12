package tests

import (
	"net/http"
)

type MockTodoController struct {
}

func (mtc *MockTodoController) GetTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (mtc *MockTodoController) PostTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (mtc *MockTodoController) PutTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (mtc *MockTodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
