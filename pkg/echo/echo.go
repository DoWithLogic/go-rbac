package echo

import (
	"net/http"

	"github.com/DoWithLogic/go-rbac/config"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// configCORS contains the CORS (Cross-Origin Resource Sharing) configuration for the server.
var configCORS = middleware.CORSConfig{
	AllowOrigins: []string{"*"},
	AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch},
}

// NewEchoServer creates and configures a new Echo server instance.
// Parameters:
//   - cfg: The application configuration.
//
// Returns:
//   - *echo.Echo: A configured Echo server instance.
func NewEchoServer(cfg config.ServerConfig) *echo.Echo {
	e := echo.New()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{DisableStackAll: true}))
	e.Use(middleware.CORSWithConfig(configCORS))
	e.Use(cacheWithRevalidation)
	// Set up the custom validator
	e.Validator = &CustomValidator{validator: validator.New()}

	e.HTTPErrorHandler = errorHandler
	e.Debug = cfg.Debug

	return e
}
