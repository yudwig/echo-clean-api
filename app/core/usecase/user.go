package usecase

import "github.com/yudwig/echo-clean-api/app/core/domain/entity/user"

type UserUseCase interface {
	CreateUser(name string) (user.User, error)
	GetUsers() ([]user.User, error)
	DeleteUser(id string) error
	UpdateUserName(id string, name string) (user.User, error)
}
