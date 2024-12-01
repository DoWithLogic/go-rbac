package usecase

import (
	"context"
	"errors"

	"github.com/DoWithLogic/go-rbac/internal/users"
	"github.com/DoWithLogic/go-rbac/internal/users/dtos"
	"github.com/DoWithLogic/go-rbac/internal/users/entities"
	"github.com/DoWithLogic/go-rbac/pkg/app_errors"
	"github.com/DoWithLogic/go-rbac/pkg/constants"
	"github.com/DoWithLogic/go-rbac/pkg/security"
	"github.com/DoWithLogic/go-rbac/pkg/utils"
	"github.com/google/uuid"
)

type Dependencies struct {
	Repositories
	Others
}

type Repositories struct {
	UserRepo users.Repositories
}

type Others struct {
	Security *security.JWTFactory
}

type usecaseImpl struct {
	userRepo users.Repositories
	security *security.JWTFactory
}

func NewUsersUC(d Dependencies) users.Usecases {
	return &usecaseImpl{
		userRepo: d.UserRepo,
		security: d.Security,
	}
}

func (uc *usecaseImpl) Login(ctx context.Context, request dtos.LoginRequest) (response dtos.LoginResponse, err error) {
	userDetail, err := uc.userRepo.Detail(ctx, entities.WithEmail(request.Email))
	if err != nil {
		if errors.Is(err, app_errors.ErrUserNotFound) {
			return response, app_errors.BadRequest(app_errors.ErrYourEmailOrPasswordWrong)
		}

		return response, err
	}

	if !utils.CheckPasswordHash(request.Password.String(), userDetail.Password) {
		return response, app_errors.BadRequest(app_errors.ErrYourEmailOrPasswordWrong)
	}

	permissions, err := uc.userRepo.RolePermissions(ctx, userDetail.RoleID)
	if err != nil {
		return response, err
	}

	response = entities.MapLoginResponse(userDetail, permissions)

	response.AccessToken, err = uc.security.CreateJWT(userDetail.ID, userDetail.RoleID, response.User.Role, permissions.GetPermissions()...)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (uc *usecaseImpl) Create(ctx context.Context, request dtos.CreateUserRequest) error {
	newUser := entities.NewCreateUser(request)
	if request.HasPassword() {
		hashedPassword, err := utils.HashPassword(request.Password.String())
		if err != nil {
			return err
		}

		newUser.SetPassword(hashedPassword)
	}

	roleID, exist := constants.MapRoleID[request.Role]
	if !exist {
		return app_errors.BadRequest(app_errors.ErrInvalidRole)
	}
	newUser.SetRoleID(uuid.MustParse(roleID.String()))

	return uc.userRepo.Create(ctx, newUser)
}
