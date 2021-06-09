package mock

import (
	"github.com/yudwig/echo-sample/app/core/domain/entity/user"
)

type UsersRepository struct {
}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{}
}

func (u UsersRepository) Get(id string) (user.User, error) {
	return *user.NewUser("id-test", "qwer"), nil
}

func (u UsersRepository) GetAll() ([]user.User, error) {
	return []user.User{*user.NewUser("id-1", "asdf"), *user.NewUser("id-2", "zxcv")}, nil
}

func (u UsersRepository) Create(name string) (user.User, error) {
	return *user.NewUser("id-test", "qwer"), nil
}

func (u UsersRepository) Delete(id string) error {
	return nil
}

func (u UsersRepository) Update(id string, name string) (user.User, error) {
	return *user.NewUser("id-test", "qwer"), nil
}
