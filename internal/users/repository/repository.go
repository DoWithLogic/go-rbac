package repository

import (
	"context"
	"errors"

	"github.com/DoWithLogic/go-rbac/internal/users"
	"github.com/DoWithLogic/go-rbac/internal/users/entities"
	"github.com/DoWithLogic/go-rbac/pkg/app_errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) users.Repositories {
	return &repositoryImpl{db}
}

func (r *repositoryImpl) Detail(ctx context.Context, opts ...entities.UserDetailOpts) (userDetail entities.UserDetail, err error) {
	args := new(entities.UserDetailRequest)
	for _, opt := range opts {
		opt(args)
	}

	query := r.db.WithContext(ctx).Table("users u").Select(`
		u.id,
		u.role_id,
		r.name AS role,
		u.username,
		u.email,
		u.password,
		u.created_at,
		u.updated_at`).
		Joins(`JOIN roles r ON r.id = u.role_id`)

	if args.ID != nil {
		query = query.Where("u.id = ?", args.ID)
	}

	if args.Email != nil {
		query = query.Where("u.email = ?", args.Email)
	}

	if err := query.Take(&userDetail).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userDetail, app_errors.NotFound(app_errors.ErrUserNotFound)
		}

		return userDetail, app_errors.InternalServerError(err)
	}

	return userDetail, nil
}

func (r *repositoryImpl) Create(ctx context.Context, request *entities.CreateUser) error {
	return r.db.WithContext(ctx).Table("users").Create(request).Error
}

func (r *repositoryImpl) RolePermissions(ctx context.Context, roleID uuid.UUID) (permissions entities.RolePermissions, err error) {
	err = r.db.WithContext(ctx).Debug().Raw(`
		SELECT
			rp.role_id,
			rp.permission_id,
			p.name AS permission
		FROM role_permissions rp 
		JOIN permissions p ON p.id = rp.permission_id
		WHERE rp.role_id = ?;
	`, roleID).Find(&permissions).Error

	if err != nil {
		return nil, err
	}

	return permissions, nil
}
