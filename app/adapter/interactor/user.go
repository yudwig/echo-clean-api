package interactor

import (
	"github.com/yudwig/echo-sample/app/adapter/repository"
)

type repositories struct {
	User repository.Users
	Log  repository.Log
}

type UserInteractor struct {
	repositories *repositories
}

func NewUserInteractor(usersRepository repository.Users, logger repository.Log) *UserInteractor {
	return &UserInteractor{
		repositories: &repositories{
			User: usersRepository,
			Log:  logger,
		},
	}
}

func (u UserInteractor) CreateUser(name string) error {
	user, err := u.repositories.User.Create(name)

	u.repositories.Log.Write(user.Name, name)

	return err
}

func (u UserInteractor) GetUsers() error {
	return nil
}

func (u UserInteractor) GetUser(id string) error {
	return nil
}

func (u UserInteractor) DeleteUser(id string) error {
	return nil
}

func (u UserInteractor) UpdateUserName(id string, name string) error {
	return nil
}
