package entities

import (
	"time"

	"github.com/DoWithLogic/go-rbac/internal/users/dtos"
	"github.com/DoWithLogic/go-rbac/pkg/constants"
	"github.com/google/uuid"
)

type UserDetailRequest struct {
	ID    *string
	Email *string
}

type UserDetailOpts func(*UserDetailRequest)

func WithID(id string) UserDetailOpts {
	return func(udr *UserDetailRequest) {
		udr.ID = &id
	}
}

func WithEmail(email string) UserDetailOpts {
	return func(udr *UserDetailRequest) {
		udr.Email = &email
	}
}

type UserDetail struct {
	ID        uuid.UUID          `gorm:"column:id"`
	RoleID    uuid.UUID          `gorm:"column:role_id"`
	Role      constants.UserRole `gorm:"column:role"`
	Username  string             `gorm:"column:username"`
	Email     string             `gorm:"column:email"`
	Password  string             `gorm:"column:password"`
	CreatedAt time.Time          `gorm:"column:created_at"`
	UpdatedAt *time.Time         `gorm:"column:updated_at"`
}

func MapLoginResponse(ud UserDetail, p RolePermissions) dtos.LoginResponse {
	return dtos.LoginResponse{
		User: dtos.LoginData{
			UserID:      ud.ID,
			Email:       ud.Email,
			RoleID:      ud.RoleID,
			Role:        ud.Role,
			Permissions: MapUserPermissions(p),
			CreatedAt:   ud.CreatedAt,
			UpdatedAt:   ud.UpdatedAt,
		},
	}
}
