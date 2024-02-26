package domain

import "todo/internal/domain/todo"

type Services struct {
	Todo *todo.TodoService
}

func NewServices() *Services {
	return &Services{
		Todo: todo.NewTodoService(),
	}
}