package controllers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"prac-go-mvc-rest/tests"
)

var mux *http.ServeMux

func TestMain(m *testing.M) {
	setUp()

	code := m.Run()
	os.Exit(code)
}

func setUp() {
	target := NewRouter(&tests.MockTodoController{})
	mux = http.NewServeMux()
	mux.HandleFunc("/", target.HandleTodosRequest)
}

func TestGetTodo(t *testing.T) {
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)
	if w.Code != 200 {
		t.Errorf("Response code is %v", w.Code)
	}
}

func TestPostTodo(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","content":"test-content"}`)
	r, _ := http.NewRequest("POST", "/", json)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 201 {
		t.Errorf("Response code is %v", w.Code)
	}
}

func TestPutTodo(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","content":"test-content"}`)
	r, _ := http.NewRequest("PUT", "/2", json)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 204 {
		t.Errorf("Response code is %v", w.Code)
	}
}

func TestDeleteTodo(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/2", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 204 {
		t.Errorf("Response code is %v", w.Code)
	}
}

func TestInvalidMethod(t *testing.T) {
	r, _ := http.NewRequest("PATCH", "/todos/", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 405 {
		t.Errorf("Method Not Allowed, Response code is %v", w.Code)
	}
}
