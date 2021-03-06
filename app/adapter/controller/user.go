package controller

import (
	"github.com/yudwig/echo-clean-api/app/adapter/interactor"
	"github.com/yudwig/echo-clean-api/app/adapter/presenter"
	"github.com/yudwig/echo-clean-api/app/adapter/repository"
	"github.com/yudwig/echo-clean-api/app/core/presentation/response"
	"github.com/yudwig/echo-clean-api/app/core/usecase"
	"github.com/yudwig/echo-clean-api/app/driver/db/rdb"
	"github.com/yudwig/echo-clean-api/app/driver/log"
)

type repositories struct {
	User repository.Users
	Log  repository.Log
}

type useCases struct {
	User usecase.UserUseCase
}

type presenters struct {
	User presenter.UserPresenter
}

type UserController struct {
	repositories *repositories
	useCases     *useCases
	presenters   *presenters
}

func NewUserController() (*UserController, error) {
	db, err := rdb.GetDbInstance()
	if err != nil {
		return nil, err
	}
	r := &repositories{
		User: rdb.NewUsersRepository(db),
		Log:  log.NewStdoutLogger(),
	}
	u := &useCases{
		User: interactor.NewUserInteractor(r.User, r.Log),
	}
	p := &presenters{
		User: presenter.NewUserPresenter(),
	}
	c := &UserController{
		repositories: r,
		useCases:     u,
		presenters:   p,
	}
	return c, nil
}

func (c UserController) CreateUser(name string) response.CreateUserResponse {
	usr, err := c.useCases.User.CreateUser(name)
	return *c.presenters.User.MakeCreateUserResponse(usr, err)
}

func (c UserController) UpdateUserName(id string, name string) response.UpdateUserResponse {
	usr, err := c.useCases.User.UpdateUserName(id, name)
	return *c.presenters.User.MakeUpdateUserResponse(usr, err)
}

func (c UserController) GetUsers() response.GetUsersResponse {
	users, err := c.useCases.User.GetUsers()
	return *c.presenters.User.MakeGetUsersResponse(users, err)
}

func (c UserController) DeleteUser(id string) response.DeleteUserResponse {
	err := c.useCases.User.DeleteUser(id)
	return *c.presenters.User.MakeDeleteUserResponse(err)
}
