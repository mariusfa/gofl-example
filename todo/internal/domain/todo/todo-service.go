package todo

type TodoServiceContract interface {
	GetTodos() []Todo
}

type TodoService struct{}

func NewTodoService() *TodoService {
	return &TodoService{}
}

func (s *TodoService) GetTodos() []Todo {
	return []Todo{
		{Title: "todo 1"},
		{Title: "todo 2"},
	}
}
