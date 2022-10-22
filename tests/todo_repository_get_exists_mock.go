package tests

import "prac-go-mvc-rest/models/entities"

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

func (mtrgex *MockTodoRepositoryGetExist) UpdateTodo(todos entities.TodoEntity) (err error) {
	return
}

func (mtrgex *MockTodoRepositoryGetExist) DeleteTodo(id int) (err error) {
	return
}
