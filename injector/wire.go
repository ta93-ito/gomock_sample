//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/ta93-ito/gomock_sample/infra"

	"github.com/google/wire"
	"github.com/ta93-ito/gomock_sample/handler"
	"github.com/ta93-ito/gomock_sample/usecase"
)

func NewUserHandler() *handler.UserHandler {
	wire.Build(handler.NewUserHandler, usecase.NewUserUsecase, infra.NewUserRepository)
	return nil
}
