package dtos

import (
	"time"

	"github.com/DoWithLogic/go-rbac/pkg/constants"
	"github.com/DoWithLogic/go-rbac/pkg/types"
	"github.com/google/uuid"
)

type LoginRequest struct {
	Email    string       `json:"email" validate:"required,email"`
	Password types.Secret `json:"password" validate:"required"`
}

type LoginResponse struct {
	User        LoginData `json:"user"`
	AccessToken string    `json:"access_token"`
}

func (lr *LoginResponse) SetPermissions(permissions UserPermissions) {
	lr.User.Permissions = permissions
}

type LoginData struct {
	UserID      uuid.UUID          `json:"id"`
	Email       string             `json:"email"`
	RoleID      uuid.UUID          `json:"role_id"`
	Role        constants.UserRole `json:"role"`
	Permissions UserPermissions    `json:"permissions"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   *time.Time         `json:"updated_at"`
}

type UserPermission struct {
	PermissionID uuid.UUID            `json:"permission_id"`
	Permission   constants.Permission `json:"permission"`
}

type UserPermissions []UserPermission
