package api

import (
	"github.com/DoWithLogic/go-rbac/internal/middlewares"
	"github.com/DoWithLogic/go-rbac/internal/users"
	"github.com/DoWithLogic/go-rbac/pkg/constants"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
)

func MapUserRoutes(g *echo.Group, h users.Handlers, mw *middlewares.Middleware) {
	users := g.Group("/users", mw.JWTMiddleware())
	users.POST("", h.CreateUserHandler, mw.RolesMiddleware(constants.AdminUserRole), mw.PermissionsMiddleware(constants.UsersCreatePermission))
}
