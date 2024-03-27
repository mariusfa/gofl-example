package todo

import (
	"testing"
	"todo/internal/database"
	"todo/internal/services/todo"

	"github.com/google/uuid"
)

func TestInsertGetTodos(t *testing.T) {
	db, err := database.Setup(&dbConfig)
	if err != nil {
		t.Fatalf("Failed to setup database: %v", err)
	}
	defer db.Close()
	todoRepository := NewTodoRepository(db)

	todoToInsert := todo.Todo{
		Id:    uuid.New(),
		Title: "Test todo",
	}

	if err := todoRepository.Insert(todoToInsert); err != nil {
		t.Fatalf("Failed to insert todo: %v", err)
	}

	todos, err := todoRepository.GetTodos()
	if err != nil {
		t.Fatalf("Failed to get todos: %v", err)
	}

	if len(todos) != 1 {
		t.Fatalf("Expected 1 todo, got %d", len(todos))
	}
	if todos[0].Id != todoToInsert.Id {
		t.Fatalf("Expected todo id to be %s, got %s", todoToInsert.Id, todos[0].Id)
	}
	if todos[0].Title != todoToInsert.Title {
		t.Fatalf("Expected todo title to be %s, got %s", todoToInsert.Title, todos[0].Title)
	}
}
