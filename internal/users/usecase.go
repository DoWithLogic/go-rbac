package users

import (
	"context"

	"github.com/DoWithLogic/go-rbac/internal/users/dtos"
)

type Usecases interface {
	Create(ctx context.Context, request dtos.CreateUserRequest) error
}
