package interactor

import (
	"github.com/yudwig/echo-clean-api/app/adapter/repository"
	"github.com/yudwig/echo-clean-api/app/core/domain/entity/user"
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

func (u UserInteractor) CreateUser(name string) (user.User, error) {
	usr, err := u.repositories.User.Create(name)
	if err != nil {
		u.repositories.Log.Write("User created failed. name: " + name)
		return user.User{}, err
	}
	u.repositories.Log.Write("User created success. user: " + usr.String())
	return usr, err
}

func (u UserInteractor) GetUsers() ([]user.User, error) {
	users, err := u.repositories.User.GetAll()
	if err != nil {
		u.repositories.Log.Write("GetUsers failed.")
		return []user.User{}, err
	}
	u.repositories.Log.Write("GetUsers success.")
	return users, err
}

func (u UserInteractor) DeleteUser(id string) error {
	err := u.repositories.User.Delete(id)
	if err != nil {
		u.repositories.Log.Write("User delete failed. id: " + id)
		return err
	}
	u.repositories.Log.Write("User delete success. id: " + id)
	return nil
}

func (u UserInteractor) UpdateUserName(id string, name string) (user.User, error) {
	usr, err := u.repositories.User.Update(id, name)
	if err != nil {
		u.repositories.Log.Write("UserName update failed. id: " + id + ", name: " + name)
		return user.User{}, err
	}
	u.repositories.Log.Write("UserName update success. user: " + usr.String())
	return usr, err
}
