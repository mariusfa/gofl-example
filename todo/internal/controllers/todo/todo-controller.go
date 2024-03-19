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
	GetTodos() []todo.Todo
}

func NewTodoController(todoService TodoService) *TodoController {
	return &TodoController{
		todoService: todoService,
	}
}

func (c *TodoController) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/todo", c.getTodos)
}

func (c *TodoController) getTodos(w http.ResponseWriter, r *http.Request) {
	listOfTodos := c.todoService.GetTodos()
	listOfDTO := fromDomainToDTOs(listOfTodos)

	jsonTodos, err := json.Marshal(listOfDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonTodos)
}
