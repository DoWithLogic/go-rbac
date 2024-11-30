package api

import (
	"github.com/DoWithLogic/go-rbac/internal/auths"
	"github.com/DoWithLogic/go-rbac/internal/middlewares"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
)

func MapAuthRoutes(group echo.Group, auths auths.Handlers, mw *middlewares.Middleware) {

}
