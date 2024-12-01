package entities

import (
	"time"

	"github.com/DoWithLogic/go-rbac/internal/users/dtos"
	"github.com/google/uuid"
)

type CreateUser struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey"`
	Username  string    `gorm:"column:username"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	RoleID    uuid.UUID `gorm:"column:role_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func NewCreateUser(request dtos.CreateUserRequest) *CreateUser {
	return &CreateUser{
		ID:        uuid.New(),
		Username:  request.Username,
		Email:     request.Username,
		CreatedAt: time.Now(),
	}
}

func (cu *CreateUser) SetRoleID(roleID uuid.UUID) {
	cu.RoleID = roleID
}

func (cu *CreateUser) SetPassword(hashedPassword string) {
	cu.Password = hashedPassword
}
