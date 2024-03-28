package todo

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo/internal/services/todo"

	"github.com/google/uuid"
)

func TestGetTodos(t *testing.T) {
	router := http.NewServeMux()
	todos := []todo.Todo{
		{Id: uuid.New(), Title: "todo 1"},
		{Id: uuid.New(), Title: "todo 2"},
	}
	fake := &TodoServiceFake{
		Todos: todos,
	}
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

	// expected := `[{"title":"todo 1"},{"title":"todo 2"}]`
	expected := `[{"id":"` + todos[0].Id.String() + `","title":"todo 1"},{"id":"` + todos[1].Id.String() + `","title":"todo 2"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestPostTodo(t *testing.T) {
	router := http.NewServeMux()
	fake := &TodoServiceFake{}
	todoController := NewTodoController(fake)
	todoController.RegisterRoutes(router)

	requestBody := strings.NewReader(`{"title":"todo 1"}`)
	req, _ := http.NewRequest("POST", "/todo", requestBody)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	if len(fake.Todos) != 1 {
		t.Errorf("handler did not insert todo")
	}
}
