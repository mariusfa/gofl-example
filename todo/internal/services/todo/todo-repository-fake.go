package todo

type fakeTodoRepository struct {
	Todos []Todo
}

func NewFakeTodoRepository() *fakeTodoRepository {
    return &fakeTodoRepository{
        Todos: []Todo{},
    }
}

func (r *fakeTodoRepository) GetTodos() ([]Todo, error) {
    return r.Todos, nil
}

func (r *fakeTodoRepository) Insert(todo Todo) error {
    r.Todos = append(r.Todos, todo)
    return nil
}

