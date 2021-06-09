package controller

import (
	"github.com/yudwig/echo-sample/app/adapter/interactor"
	"github.com/yudwig/echo-sample/app/adapter/repository"
	"github.com/yudwig/echo-sample/app/core/usecase"
	"github.com/yudwig/echo-sample/app/driver/db/mock"
	"github.com/yudwig/echo-sample/app/driver/log"
)

type repositories struct {
	User repository.Users
	Log  repository.Log
}

type useCases struct {
	User usecase.UserUseCase
}

type UserController struct {
	repositories *repositories
	useCases     *useCases
}

func NewUserController() *UserController {

	r := &repositories{
		User: mock.NewUsersRepository(),
		Log:  log.NewStdoutLogger(),
	}

	u := &useCases{
		User: interactor.NewUserInteractor(r.User, r.Log),
	}

	c := &UserController{
		repositories: r,
		useCases:     u,
	}

	return c
}

func (c UserController) CreateUser(name string) error {
	return c.useCases.User.CreateUser(name)
}
