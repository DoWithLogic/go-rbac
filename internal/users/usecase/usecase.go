package usecase

import (
	"context"

	"github.com/DoWithLogic/go-rbac/internal/users"
	"github.com/DoWithLogic/go-rbac/internal/users/dtos"
	"github.com/DoWithLogic/go-rbac/internal/users/entities"
	"github.com/DoWithLogic/go-rbac/pkg/app_errors"
	"github.com/DoWithLogic/go-rbac/pkg/constants"
	"github.com/DoWithLogic/go-rbac/pkg/utils"
	"github.com/google/uuid"
)

type Dependencies struct {
	Repositories
}

type Repositories struct {
	UserRepo users.Repositories
}

type usecaseImpl struct {
	userRepo users.Repositories
}

func NewUsersUC(d Dependencies) users.Usecases {
	return &usecaseImpl{
		userRepo: d.UserRepo,
	}
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
