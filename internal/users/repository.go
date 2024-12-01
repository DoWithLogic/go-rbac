package users

import (
	"context"

	"github.com/DoWithLogic/go-rbac/internal/users/entities"
)

type Repositories interface {
	Create(ctx context.Context, request *entities.CreateUser) error
}
