package controllers

import (
	"encoding/json"
	"net/http/httptest"
	"prac-go-mvc-rest/controllers/dto"
	"prac-go-mvc-rest/tests"
	"strings"
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

func TestPostTodo_Exist(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","content":"test-content"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", json)

	target := NewTodoController(&tests.MockTodoRepository{})
	target.PostTodo(w, r)

	if w.Code != 201 {
		t.Errorf("Response code is %v", w.Code)
	}
	if w.Header().Get("Location") != r.Host+r.URL.Path+"2" {
		t.Errorf("Location is %v", w.Header().Get("Location"))
	}
}

func TestPostTodo_Error(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","contents":"test-content"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", json)

	target := NewTodoController(&tests.MockTodoRepositoryError{})
	target.PostTodo(w, r)

	if w.Code != 500 {
		t.Errorf("Response code is %v", w.Code)
	}
	if w.Header().Get("Location") != "" {
		t.Errorf("Location is %v", w.Header().Get("Location"))
	}
}

func TestPutTodo_Exist(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","contents":"test-content"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/2", json)

	target := NewTodoController(&tests.MockTodoRepository{})
	target.PutTodo(w, r)

	if w.Code != 204 {
		t.Errorf("Response code is %v", w.Code)
	}
}

func TestPutTodo_InvalidPath(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","contents":"test-content"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/", json)

	target := NewTodoController(&tests.MockTodoRepository{})
	target.PutTodo(w, r)

	if w.Code != 400 {
		t.Errorf("Response code %v", w.Code)
	}
}

func TestPutTodo_Error(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","contents":"test-content"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/2", json)

	target := NewTodoController(&tests.MockTodoRepositoryError{})
	target.PutTodo(w, r)

	if w.Code != 500 {
		t.Errorf("Response code %v", w.Code)
	}
}

func TestDelete_Exist(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/2", nil)

	target := NewTodoController(&tests.MockTodoRepository{})
	target.DeleteTodo(w, r)

	if w.Code != 204 {
		t.Errorf("Response code is %v", w.Code)
	}
}

func TestDelete_InvalidPath(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/", nil)

	target := NewTodoController(&tests.MockTodoRepository{})
	target.DeleteTodo(w, r)

	if w.Code != 400 {
		t.Errorf("Response code is %v", w.Code)
	}
}

func TestDelete_Error(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/2", nil)

	target := NewTodoController(&tests.MockTodoRepositoryError{})
	target.DeleteTodo(w, r)

	if w.Code != 500 {
		t.Errorf("Response code is %v", w.Code)
	}
}
