package todo

import "github.com/google/uuid"

type Todo struct {
	Id    uuid.UUID
	Title string
}
