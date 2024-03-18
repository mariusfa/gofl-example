package internal

import "todo/internal/todo"

type Services struct {
	Todo *todo.TodoService
}

func NewServices() *Services {
	return &Services{
		Todo: todo.NewTodoService(),
	}
}
