package app

import (
	"encoding/json"
	"net/http"
)

type TodoController struct{}

func NewTodoController() *TodoController {
	return &TodoController{}
}

func (c *TodoController) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/todo", c.getTodos)
}

func (c *TodoController) getTodos(w http.ResponseWriter, r *http.Request) {
	listOfTodos := []TodoDTO{
		{Title: "todo 1"},
		{Title: "todo 2"},
	}

	jsonTodos, err := json.Marshal(listOfTodos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonTodos)
}
