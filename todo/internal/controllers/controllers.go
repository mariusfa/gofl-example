package controllers

import (
	"net/http"
	"todo/internal/controllers/todo"
	"todo/internal/services"

	hc "github.com/mariusfa/gofl/v2/health-controller"
	mc "github.com/mariusfa/gofl/v2/metrics-controller"
)

type Controllers struct {
	health  *hc.HealthController
	metrics *mc.MetricsController
	todo    *todo.TodoController
}

func New(services *services.Services) *Controllers {
	return &Controllers{
		health:  hc.NewHealthController(),
		metrics: mc.NewMetricsController(),
		todo:    todo.NewTodoController(services.Todo),
	}
}

func (c *Controllers) RegisterRoutes(router *http.ServeMux) {
	c.health.RegisterRoutes(router)
	c.metrics.RegisterRoutes(router)
	c.todo.RegisterRoutes(router)
}
