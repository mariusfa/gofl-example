package todo

import "todo/internal/services/todo"

type TodoDTO struct {
	Id    string `json:"id"`
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
		Id:    todo.Id.String(),
		Title: todo.Title,
	}
}

type TodoRequestDTO struct {
	Title string `json:"title"`
}
