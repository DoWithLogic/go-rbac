package users

import (
	"context"

	"github.com/DoWithLogic/go-rbac/internal/users/entities"
	"github.com/google/uuid"
)

type Repositories interface {
	Detail(ctx context.Context, opts ...entities.UserDetailOpts) (userDetail entities.UserDetail, err error)
	Create(ctx context.Context, request *entities.CreateUser) error
	RolePermissions(ctx context.Context, roleID uuid.UUID) (permissions entities.RolePermissions, err error)
}
