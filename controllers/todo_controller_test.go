package controllers

import (
	"encoding/json"
	"net/http/httptest"
	"prac-go-mvc-rest/controllers/dto"
	"prac-go-mvc-rest/tests"
	"testing"
)

func TestGetTodo_NotFound(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	target := NewTodoController(&tests.MockTodoRepository{})
	target.GetTodo(w, r)

	if w.Code != 200 {
		t.Errorf("Response code is %v", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Content-Type is %v", w.Header().Get("Content-Type"))
	}

	body := make([]byte, w.Body.Len())
	w.Body.Read(body)
	var todosResponse dto.TodosResponse
	json.Unmarshal(body, &todosResponse)
	if len(todosResponse.Todos) != 0 {
		t.Errorf("Response is %v", todosResponse)
	}
}

func TestGetTodo_Exist(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	target := NewTodoController(&tests.MockTodoRepositoryGetExist{})
	target.GetTodo(w, r)

	if w.Code != 200 {
		t.Errorf("Response code is %v", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Content-Type is %v", w.Header().Get("Content-Type"))
	}

	body := make([]byte, w.Body.Len())
	w.Body.Read(body)
	var todosResponse dto.TodosResponse
	json.Unmarshal(body, &todosResponse.Todos)
	if len(todosResponse.Todos) != 2 {
		t.Errorf("Response is %v", todosResponse)
	}
}

func TestGetTodo_Error(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	target := NewTodoController(&tests.MockTodoRepositoryError{})
	target.GetTodo(w, r)

	if w.Code != 500 {
		t.Errorf("Response code is %v", w.Code)
	}
	if w.Header().Get("Content-Type") != "" {
		t.Errorf("Content-Type is %v", w.Header().Get("Content-Type"))
	}

	if w.Body.Len() != 0 {
		t.Errorf("body is %v", w.Body.Len())
	}
}
