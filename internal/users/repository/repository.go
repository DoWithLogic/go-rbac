package repository

import (
	"context"

	"github.com/DoWithLogic/go-rbac/internal/users"
	"github.com/DoWithLogic/go-rbac/internal/users/entities"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) users.Repositories {
	return &repositoryImpl{db}
}

func (r *repositoryImpl) Create(ctx context.Context, request *entities.CreateUser) error {
	return r.db.WithContext(ctx).Table("users").Create(request).Error
}
