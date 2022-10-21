package repositories

import (
	"log"
	"prac-go-mvc-rest/models/entities"
)

type TodoRepository interface {
	GetTodo() (todos []entities.TodoEntity, err error)
	InsertTodo(todo entities.TodoEntity) (id int, err error)
	UpdateTodo(todo entities.TodoEntity) (err error)
}

type todoRepository struct {
}

func NewTodoRepository() TodoRepository {
	return &todoRepository{}
}

func (tr *todoRepository) GetTodo() (todos []entities.TodoEntity, err error) {
	todos = []entities.TodoEntity{}
	rows, err := Db.Query("SELECT id, title, content FROM todo ORDER BY id DESC")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		todo := entities.TodoEntity{}
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Content)
		if err != nil {
			log.Print(err)
		}
		todos = append(todos, todo)
	}

	return
}

func (tr *todoRepository) InsertTodo(todo entities.TodoEntity) (id int, err error) {
	_, err = Db.Exec("INSERT INTO todo (title, content) VALUE (?, ?)", todo.Title, todo.Content)
	if err != nil {
		log.Print(err)
		return
	}

	err = Db.QueryRow("SELECT id FROM todo ORDER BY id DESC LIMIT 1").Scan(&id)

	return
}

func (tr *todoRepository) UpdateTodo(todo entities.TodoEntity) (err error) {
	_, err = Db.Exec("UPDATE todo SET title = ?, content = ? WHERE id = ?", todo.Title, todo.Content, todo.Id)
	return
}
