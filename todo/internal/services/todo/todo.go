package todo

import "github.com/google/uuid"

type Todo struct {
	Id    uuid.UUID
	Title string
}

func NewTodo(title string) Todo {
	return Todo{
		Id:    uuid.New(),
		Title: title,
	}
}
