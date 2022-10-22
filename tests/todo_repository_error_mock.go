package tests

import (
	"errors"
	"prac-go-mvc-rest/models/entities"
)

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

func (mtrie *MockTodoRepositoryError) UpdateTodo(todo entities.TodoEntity) (err error) {
	err = errors.New("unexpected error occurred")
	return
}

func (mtrie *MockTodoRepositoryError) DeleteTodo(id int) (err error) {
	err = errors.New("unexpected error occurred")
	return
}
