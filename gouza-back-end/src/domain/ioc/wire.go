//go:build wireinject
// +build wireinject

package ioc

import (
	"github.com/google/wire"
	"gouza-back-end/src/action"
	"gouza-back-end/src/domain/repo"
	"gouza-back-end/src/domain/service"
)

func BuildUserAction() (*action.UserAction, error) {
	wire.Build(action.NewUserAction, service.NewUserService, repo.NewUserRepository, repo.NewDbConnection)
	return new(action.UserAction), nil
}

func BuildUserService() (*service.UserService, error) {
	wire.Build(service.NewUserService, repo.NewUserRepository, repo.NewDbConnection)
	return new(service.UserService), nil
}
