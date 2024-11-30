package usecase

import "github.com/DoWithLogic/go-rbac/internal/auths"

type Dependencies struct {
	Repositories
}

type Repositories struct {
	AuthRepo auths.Repositories
}

type usecaseImpl struct {
	authRepo auths.Repositories
}

func NewAuthUC(d Dependencies) auths.Usecases {
	return &usecaseImpl{
		authRepo: d.AuthRepo,
	}
}
