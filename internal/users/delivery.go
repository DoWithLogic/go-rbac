package users

import "github.com/labstack/echo/v4"

type Handlers interface {
	CreateUserHandler(c echo.Context) error
	LoginHandler(c echo.Context) error
}
