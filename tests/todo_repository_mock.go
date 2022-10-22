package tests

import "prac-go-mvc-rest/models/entities"

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

func (mtr *MockTodoRepository) UpdateTodo(todo entities.TodoEntity) (err error) {
	return
}

func (mtr *MockTodoRepository) DeleteTodo(id int) (err error) {
	return
}
