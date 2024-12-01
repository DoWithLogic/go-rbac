package users

import (
	"context"

	"github.com/DoWithLogic/go-rbac/internal/users/dtos"
)

type Usecases interface {
	Login(ctx context.Context, request dtos.LoginRequest) (response dtos.LoginResponse, err error)
	Create(ctx context.Context, request dtos.CreateUserRequest) error
}
