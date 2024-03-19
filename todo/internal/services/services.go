package services

import "todo/internal/services/todo"

type Services struct {
	Todo *todo.TodoService
}

func New() *Services {
	return &Services{
		Todo: todo.NewTodoService(),
	}
}
