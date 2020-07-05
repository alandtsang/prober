package metirics

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// swagger:route GET /metrics Metrics metricsHandler
//
// Metrics handler
//
// Responses:
//   200: response
func MetricsHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "metrics")
}
