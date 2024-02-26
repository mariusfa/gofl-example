package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/internal/domain/todo"
)

type TodoServiceFake struct{}

func (s *TodoServiceFake) GetTodos() []todo.Todo {
	return []todo.Todo{
		{Title: "todo 1"},
		{Title: "todo 2"},
	}
}

func TestGetTodos(t *testing.T) {
	router := http.NewServeMux()
	fake := &TodoServiceFake{}
	todoController := NewTodoController(fake)
	todoController.RegisterRoutes(router)

	req, err := http.NewRequest("GET", "/todo", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"title":"todo 1"},{"title":"todo 2"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
