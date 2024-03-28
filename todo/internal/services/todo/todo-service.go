package todo

type TodoService struct {
	todoRepository TodoRepository
}

type TodoRepository interface {
	GetTodos() ([]Todo, error)
	Insert(todo Todo) error
}

func NewTodoService(todoRepository TodoRepository) *TodoService {
	return &TodoService{
		todoRepository: todoRepository,
	}
}

func (s *TodoService) GetTodos() ([]Todo, error) {
	return s.todoRepository.GetTodos()
}

func (s *TodoService) Insert(todo Todo) error {
	return s.todoRepository.Insert(todo)
}
