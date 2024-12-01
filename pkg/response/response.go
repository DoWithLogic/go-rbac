package response

import (
	"errors"
	"net/http"

	"github.com/DoWithLogic/go-rbac/pkg/app_errors"
	"github.com/DoWithLogic/go-rbac/pkg/constants"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// FailedResponse represents a failed response structure for API responses.
type FailedResponse struct {
	Code    int                       `json:"code" example:"500"`                      // HTTP status code.
	Message constants.ResponseMessage `json:"message" example:"internal_server_error"` // Message corresponding to the status code.
	Error   string                    `json:"error" example:"{$err}"`                  // error message.
}

// BasicResponse represents a failed response structure for API responses.
type BasicResponse struct {
	Code    int         `json:"code" example:"500"`                      // HTTP status code.
	Message string      `json:"message" example:"internal_server_error"` // Message corresponding to the status code.
	Error   string      `json:"error" example:"{$err}"`                  // error message.
	Data    interface{} `json:"data,omitempty"`
}

// BasicBuilder constructs a BasicBuilder based on the provided error.
func BasicBuilder(result BasicResponse) BasicResponse {
	return result
}

// Send sends the CustomResponse as a JSON response using the provided Echo context.
func (c BasicResponse) Send(ctx echo.Context) error {
	trace.SpanFromContext(ctx.Request().Context()).SetStatus(codes.Ok, http.StatusText(c.Code))
	return ctx.JSON(c.Code, c)
}

// ErrorBuilder constructs a FailedResponse based on the provided error.
func ErrorBuilder(err error) FailedResponse {
	var appErr *app_errors.AppError
	if errors.As(err, &appErr) {
		ae := err.(*app_errors.AppError)

		return FailedResponse{
			Code:    ae.Code,
			Message: ae.Message,
			Error:   ae.Error(),
		}
	}

	var errStr = constants.InternalServerError.String()
	if err != nil {
		errStr = err.Error()
	}

	return FailedResponse{
		Code:    http.StatusInternalServerError,
		Message: constants.InternalServerError,
		Error:   errStr,
	}
}

// Send sends the CustomResponse as a JSON response using the provided Echo context.
func (x FailedResponse) Send(c echo.Context) error {
	return c.JSON(x.Code, x)
}

// SuccessResponse represents a success response structure for API responses.
type SuccessResponse struct {
	Success
	Meta
}

type ResponseFormat struct {
	Code    int                       `json:"code" example:"200"` // HTTP status code.
	Message constants.ResponseMessage `json:"message" example:"success"`
}

type Success struct {
	ResponseFormat
	Data interface{} `json:"data,omitempty"` // data payload.
}

type Meta struct {
	Meta interface{} `json:"meta,omitempty"` //pagination payload.
	Success
}

// SuccessBuilder constructs a CustomResponse with a Success status and the provided response data.
func SuccessBuilder(response interface{}, meta ...interface{}) SuccessResponse {
	result := SuccessResponse{
		Success: Success{
			ResponseFormat: ResponseFormat{
				Code:    http.StatusOK,
				Message: constants.Success,
			},
			Data: response,
		},
	}

	if len(meta) > 0 {
		result.Meta.Meta = meta[0]
	}

	return result
}

func SuccessNoResponseBuilder() SuccessResponse {
	result := SuccessResponse{
		Success: Success{
			ResponseFormat: ResponseFormat{
				Code:    http.StatusNoContent,
				Message: constants.Success,
			},
		},
	}

	return result
}

// Send sends the CustomResponse as a JSON response using the provided Echo context.
func (c SuccessResponse) Send(ctx echo.Context) error {
	trace.SpanFromContext(ctx.Request().Context()).SetStatus(codes.Ok, http.StatusText(c.Code))
	return ctx.JSON(c.Code, c)
}
