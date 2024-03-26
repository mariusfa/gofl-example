package services

import (
	"todo/internal/repositories"
	"todo/internal/services/todo"
)

type Services struct {
	Todo *todo.TodoService
}

func New(repositories *repositories.Repositories) *Services {
	return &Services{
		Todo: todo.NewTodoService(repositories.Todo),
	}
}
