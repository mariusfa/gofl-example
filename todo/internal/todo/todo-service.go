package todo

import "github.com/google/uuid"

type TodoServiceContract interface {
	GetTodos() []Todo
}

type TodoService struct{}

func NewTodoService() *TodoService {
	return &TodoService{}
}

func (s *TodoService) GetTodos() []Todo {
	return []Todo{
		{Id: uuid.New(), Title: "todo 1"},
		{Id: uuid.New(), Title: "todo 2"},
	}
}
