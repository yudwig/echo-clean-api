package mockdb

import (
	"github.com/yudwig/echo-sample/app/core/domain/entity/user"
)

type UsersRepository struct {
}

func NewUsersRepository() UsersRepository {
	return UsersRepository{}
}

func (u UsersRepository) GetAll() ([]user.User, error) {
	usr1, _ := user.NewUser("id-1", "user 1")
	usr2, _ := user.NewUser("id-2", "user 2")
	usr3, _ := user.NewUser("id-3", "user 3")
	return []user.User{usr1, usr2, usr3}, nil
}

func (u UsersRepository) Create(name string) (user.User, error) {
	usr, _ := user.NewUser("id-1", name)
	return usr, nil
}

func (u UsersRepository) Delete(id string) error {
	return nil
}

func (u UsersRepository) Update(id string, name string) (user.User, error) {
	usr, _ := user.NewUser(id, name)
	return usr, nil
}
