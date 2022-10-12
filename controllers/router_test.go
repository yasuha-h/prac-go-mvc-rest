package controllers

import (
	"net/http"
	"net/http/httptest"
	"os"
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
