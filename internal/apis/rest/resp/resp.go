package resp

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/alandtsang/prober/internal/errs"
)

// swagger:response response
type ResponseWrapper struct {
	// in:body
	Body Response `json:"body"`
}

type Response struct {
	Code     errs.StatusCode `json:"code"`
	Message  string          `json:"message"`
	Response interface{}     `json:"response"`
}

// BuildBadRequestErrorResponse builds bad request error response.
func BuildBadRequestErrorResponse(c echo.Context, err error) error {
	return BuildErrorMessageResponse(c, http.StatusBadRequest, errs.StatusBadRequest, err.Error())
}

// BuildInternalServerErrorResponse builds internal server error response.
func BuildInternalServerErrorResponse(c echo.Context, err error) error {
	return BuildErrorMessageResponse(c, http.StatusInternalServerError, errs.StatusInternalServerError, err.Error())
}

// BuildErrorMessageResponse builds error response with message.
func BuildErrorMessageResponse(c echo.Context, httpCode int, errorCode errs.StatusCode, errorMsg string) error {
	return c.JSON(httpCode, &Response{
		Code:    errorCode,
		Message: errorMsg,
	})
}

// BuildSuccessMessageResponse builds success message response.
func BuildSuccessMessageResponse(c echo.Context, data interface{}) error {
	if data == nil {
		data = http.StatusText(http.StatusOK)
	}

	return c.JSON(http.StatusOK, &Response{
		Code:     errs.StatusOK,
		Message:  "success",
		Response: data,
	})
}
