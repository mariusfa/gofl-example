package app

import (
	"net/http"
	"todo/internal"

	hc "github.com/mariusfa/gofl/v2/health-controller"
)

type Controllers struct {
	health *hc.HealthController
	todo   *TodoController
}

func NewControllers(services *internal.Services) *Controllers {
	return &Controllers{
		health: hc.NewHealthController(),
		todo:   NewTodoController(services.Todo),
	}
}

func (c *Controllers) RegisterRoutes(router *http.ServeMux) {
	c.health.RegisterRoutes(router)
	c.todo.RegisterRoutes(router)
}
