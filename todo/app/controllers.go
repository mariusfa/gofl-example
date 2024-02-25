package app

import (
	"net/http"

	hc "github.com/mariusfa/gofl/health-controller"
)

type Controllers struct {
	health *hc.HealthController
}

func NewControllers() *Controllers {
	return &Controllers{
		health: hc.NewHealthController(),
	}
}

func (c *Controllers) RegisterRoutes(router *http.ServeMux) {
	c.health.RegisterRoutes(router)
}
