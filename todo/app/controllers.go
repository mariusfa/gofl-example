package app

import "net/http"

type Controllers struct {
	health *HealthController
}

func NewControllers() *Controllers {
	return &Controllers{
		health: NewHealthController(),
	}
}

func (c *Controllers) RegisterRoutes(router *http.ServeMux) {
	c.health.registerRoutes(router)
}
