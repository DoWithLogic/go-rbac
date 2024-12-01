package api

import (
	"github.com/DoWithLogic/go-rbac/internal/users"
	"github.com/DoWithLogic/go-rbac/internal/users/dtos"
	"github.com/DoWithLogic/go-rbac/pkg/app_errors"
	"github.com/DoWithLogic/go-rbac/pkg/response"
	"github.com/labstack/echo/v4"
)

type handlerImpl struct {
	uc users.Usecases
}

func NewUsersHandler(uc users.Usecases) users.Handlers {
	return &handlerImpl{uc: uc}
}

//	@Summary		Create User
//	@Description	Create User.
//	@ID				create-user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			body	body		dtos.CreateUserRequest	true	"User Object"
//	@Success		200		{object}	response.ResponseFormat	"SUCCESS"
//	@Failure		500		{object}	response.FailedResponse	"INTERNAL_SERVER__ERROR"
//	@Router			/users [post]
//	@Security		BearerToken
func (h *handlerImpl) CreateUserHandler(c echo.Context) error {
	var request dtos.CreateUserRequest
	if err := c.Bind(&request); err != nil {
		return response.ErrorBuilder(app_errors.BadRequest(err)).Send(c)
	}

	if err := c.Validate(request); err != nil {
		return response.ErrorBuilder(app_errors.BadRequest(err)).Send(c)
	}

	if err := h.uc.Create(c.Request().Context(), request); err != nil {
		return response.ErrorBuilder(err).Send(c)
	}

	return response.SuccessBuilder(nil).Send(c)
}
