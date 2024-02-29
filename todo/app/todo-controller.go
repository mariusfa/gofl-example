package app

import (
	"encoding/json"
	"net/http"
	"todo/internal/domain/todo"
)

type TodoController struct {
	TodoService todo.TodoServiceContract
}

func NewTodoController(todoService todo.TodoServiceContract) *TodoController {
	return &TodoController{
		TodoService: todoService,
	}
}

func (c *TodoController) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/todo", c.getTodos)
}

func (c *TodoController) getTodos(w http.ResponseWriter, r *http.Request) {
	listOfTodos := c.TodoService.GetTodos()
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
