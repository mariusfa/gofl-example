package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/internal/services"
	"todo/internal/services/todo"
)

type todoRepoFake struct{}

func (t *todoRepoFake) GetTodos() ([]todo.Todo, error) {
	return []todo.Todo{}, nil
}

func TestHealth(t *testing.T) {
	todoFake := &todoRepoFake{}
	todoService := todo.NewTodoService(todoFake)
	services := services.Services{
		Todo: todoService,
	}

	router := setup(&services)
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "OK"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
