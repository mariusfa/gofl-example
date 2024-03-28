package todo

import "todo/internal/services/todo"

type TodoServiceFake struct {
	Todos []todo.Todo
}

func (s *TodoServiceFake) GetTodos() ([]todo.Todo, error) {
	return s.Todos, nil
}

func (s *TodoServiceFake) Insert(todo todo.Todo) error {
	s.Todos = append(s.Todos, todo)
	return nil
}
