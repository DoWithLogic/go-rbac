package dtos

import (
	"github.com/DoWithLogic/go-rbac/pkg/constants"
	"github.com/DoWithLogic/go-rbac/pkg/types"
)

type CreateUserRequest struct {
	Username string             `json:"username" validate:"required"`
	Email    string             `json:"email" validate:"required,email"`
	Password types.Secret       `json:"password" validate:"required"`
	Role     constants.UserRole `json:"role" validate:"required"`
}

func (cur CreateUserRequest) HasPassword() bool {
	return cur.Password != ""
}
