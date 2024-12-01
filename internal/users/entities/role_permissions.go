package entities

import (
	"github.com/DoWithLogic/go-rbac/internal/users/dtos"
	"github.com/DoWithLogic/go-rbac/pkg/constants"
	"github.com/google/uuid"
)

type RolePermission struct {
	RoleID       uuid.UUID            `gorm:"column:role_id"`
	PermissionID uuid.UUID            `gorm:"column:permission_id"`
	Permission   constants.Permission `gorm:"column:permission"`
}

type RolePermissions []RolePermission

func MapUserPermissions(rp RolePermissions) dtos.UserPermissions {
	permissions := make(dtos.UserPermissions, 0)
	for _, v := range rp {
		permissions = append(permissions, dtos.UserPermission{
			PermissionID: v.PermissionID,
			Permission:   v.Permission,
		})
	}

	return permissions
}

func (rp RolePermissions) GetPermissions() []constants.Permission {
	permissions := make([]constants.Permission, 0)
	for _, permission := range rp {
		permissions = append(permissions, permission.Permission)
	}

	return permissions
}
