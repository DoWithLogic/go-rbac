package repository

import (
	"github.com/DoWithLogic/go-rbac/internal/auths"
	"github.com/jinzhu/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auths.Repositories {
	return &repositoryImpl{db}
}
