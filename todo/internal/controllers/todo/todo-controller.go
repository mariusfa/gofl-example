package todo

import (
	"encoding/json"
	"net/http"
	"todo/internal/services/todo"
)

type TodoController struct {
	todoService TodoService
}

type TodoService interface {
	GetTodos() ([]todo.Todo, error)
	Insert(todo.Todo) error
}

func NewTodoController(todoService TodoService) *TodoController {
	return &TodoController{
		todoService: todoService,
	}
}

func (c *TodoController) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /todo", c.getTodos)
}

func (c *TodoController) getTodos(w http.ResponseWriter, r *http.Request) {
	listOfTodos, err := c.todoService.GetTodos()
	if err != nil {
		// TODO: log error instead of returning error to client
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	listOfDTO := fromDomainToDTOs(listOfTodos)

	jsonTodos, err := json.Marshal(listOfDTO)
	if err != nil {
		// TODO: log error instead of returning error to client
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonTodos)
}
