package repositories

import (
	"database/sql"
	"todo/internal/repositories/todo"
)

type Repositories struct {
	Todo *todo.TodoRepository
}

func New(db *sql.DB) *Repositories {
	return &Repositories{
		Todo: todo.NewTodoRepository(db),
	}
}
