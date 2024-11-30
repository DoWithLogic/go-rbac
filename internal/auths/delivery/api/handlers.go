package api

import "github.com/DoWithLogic/go-rbac/internal/auths"

type handlerImpl struct {
	uc auths.Usecases
}

func NewAuthHandler(uc auths.Usecases) auths.Handlers {
	return &handlerImpl{uc: uc}
}
