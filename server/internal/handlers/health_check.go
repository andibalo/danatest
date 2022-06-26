package handlers

import (
	"net/http"
	"streaming/internal/constants"

	"github.com/labstack/echo/v4"
)

// HealthCheck is a standard, simple health check
type HealthCheck struct{}

// AddRoutes adds the routers for this API to the provided router (or subrouter)
func (h *HealthCheck) AddRoutes(e *echo.Echo) {
	e.GET(constants.HealthAPIPath, h.handler)
}

func (h *HealthCheck) handler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
