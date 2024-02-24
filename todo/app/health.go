package app

import "net/http"

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (h *HealthController) registerRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /health", h.getHealth)
}

func (h *HealthController) getHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}