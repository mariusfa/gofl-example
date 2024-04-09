package todo

import "testing"

func TestInsert(t *testing.T) {
	fakeTodoRepository := NewFakeTodoRepository()
	todoService := NewTodoService(fakeTodoRepository)

	todoToInsert := NewTodo("Test todo")
	if err := todoService.Insert(todoToInsert); err != nil {
		t.Fatalf("Failed to insert todo: %v", err)
	}

	insertedTodos := fakeTodoRepository.Todos
	if len(insertedTodos) != 1 {
		t.Fatalf("Expected 1 todo, got %d", len(insertedTodos))
	}

	if insertedTodos[0].Title != todoToInsert.Title {
		t.Fatalf("Expected todo title to be %s, got %s", todoToInsert.Title, insertedTodos[0].Title)
	}
}

func TestGetTodos(t *testing.T) {
	fakeTodoRepository := NewFakeTodoRepository()
	fakeTodoRepository.Todos = []Todo{
		NewTodo("Test todo 1"),
		NewTodo("Test todo 2"),
	}
	todoService := NewTodoService(fakeTodoRepository)

	todos, err := todoService.GetTodos()
    if err != nil {
        t.Fatalf("Failed to get todos: %v", err)
    }
	if len(todos) != 2 {
		t.Fatalf("Expected 2 todos, got %d", len(todos))
	}
	if todos[0].Title != "Test todo 1" {
		t.Fatalf("Expected todo title to be 'Test todo 1', got %s", todos[0].Title)
	}
}
