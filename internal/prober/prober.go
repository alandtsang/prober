package prober

import (
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/alandtsang/prober/internal/apis/rest/resp"
)

type probeFn func(target string) bool

var (
	probers = map[string]probeFn{
		"http": ProbeHTTP,
		//"tcp":  ProbeTCP,
		//"icmp": ProbeICMP,
		//"dns":  ProbeDNS,
	}
)

// swagger:parameters probeHandler
type probeReqWrapper struct {
	// in:body
	Params probeParams `json:"params"`
}

type probeParams struct {
	Module string `json:"module" query:"module" validate:"required"`
	Target string `json:"target" query:"modules" validate:"required"`
}

// swagger:route POST /probe Probe probeHandler
//
// Probe handler
//
// Responses:
//   200: response
func ProbeHandler(c echo.Context) error {
	var params probeParams
	if err := c.Bind(&params); err != nil {
		return resp.BuildBadRequestErrorResponse(c, err)
	}
	fmt.Printf("params=%+v\n", params)

	if err := c.Validate(&params); err != nil {
		return resp.BuildInternalServerErrorResponse(c, err)
	}

	probeFn, ok := probers[params.Module]
	if !ok {
		return resp.BuildBadRequestErrorResponse(c, errors.New("xxx"))
	}

	probeFn(params.Target)

	return resp.BuildSuccessMessageResponse(c, nil)
}
