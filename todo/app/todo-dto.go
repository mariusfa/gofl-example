package app

import "todo/internal/domain/todo"

type TodoDTO struct {
	Title string `json:"title"`
}

func fromDomainToDTOs(todos []todo.Todo) []TodoDTO {
	todoDTOs := make([]TodoDTO, len(todos))
	for i, todo := range todos {
		todoDTOs[i] = fromDomainToDTO(todo)
	}
	return todoDTOs
}

func fromDomainToDTO(todo todo.Todo) TodoDTO {
	return TodoDTO{
		Title: todo.Title,
	}
}