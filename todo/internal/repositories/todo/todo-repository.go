package todo

import (
	"database/sql"
	"todo/internal/services/todo"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) Insert(todo todo.Todo) error {
	_, err := r.db.Exec("INSERT INTO todoschema.todos (id, title) VALUES ($1, $2)", todo.Id, todo.Title)
	return err
}

func (r *TodoRepository) GetTodos() ([]todo.Todo, error) {
	rows, err := r.db.Query("SELECT id, title FROM todoschema.todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []todo.Todo
	for rows.Next() {
		var todo todo.Todo
		err := rows.Scan(&todo.Id, &todo.Title)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
