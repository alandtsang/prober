package gateway

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/alandtsang/prober/internal/metirics"
	"github.com/alandtsang/prober/internal/prober"
)

type gatewayValidator struct {
	validator *validator.Validate
}

func (gv *gatewayValidator) Validate(i interface{}) error {
	return gv.validator.Struct(i)
}

type gateway struct {
	e *echo.Echo
}

// New returns a new gateway instance.
func New() *gateway {
	e := echo.New()
	e.Validator = &gatewayValidator{validator: validator.New()}
	return &gateway{
		e: e,
	}
}

// SetupRoute sets up e.
func (g *gateway) SetupRoute() *gateway {
	// Middleware
	g.setMiddleware()

	// Routes
	g.setRoute()

	return g
}

func (gate *gateway) setMiddleware() {
	gate.e.Use(middleware.Logger())
	gate.e.Use(middleware.Recover())
}

func (g *gateway) setRoute() {
	g.e.GET("/swagger/*", echoSwagger.WrapHandler)
	g.e.GET("/metrics", metirics.MetricsHandler)
	g.e.POST("/probe", prober.ProbeHandler)
}

// ListenAndServe start listening port and serve.
func (g *gateway) ListenAndServe(port int) error {
	addr := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return g.e.StartServer(server)
}
