package controllers

import (
	"net/http"
	"todo/internal/controllers/todo"
	"todo/internal/services"

	hc "github.com/mariusfa/gofl/v2/health-controller"
)

type Controllers struct {
	health *hc.HealthController
	todo   *todo.TodoController
}

func New(services *services.Services) *Controllers {
	return &Controllers{
		health: hc.NewHealthController(),
		todo:   todo.NewTodoController(services.Todo),
	}
}

func (c *Controllers) RegisterRoutes(router *http.ServeMux) {
	c.health.RegisterRoutes(router)
	c.todo.RegisterRoutes(router)
}
