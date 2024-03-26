package todo

import "github.com/google/uuid"

type TodoService struct {
	todoRepository TodoRepository
}

type TodoRepository interface {
	GetTodos() ([]Todo, error)
}

func NewTodoService(TodoRepository TodoRepository) *TodoService {
	return &TodoService{}
}

func (s *TodoService) GetTodos() []Todo {
	return []Todo{
		{Id: uuid.New(), Title: "todo 1"},
		{Id: uuid.New(), Title: "todo 2"},
	}
}
