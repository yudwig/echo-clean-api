package rdb

import (
	"github.com/yudwig/echo-sample/app/core/domain/entity/user"
	"github.com/yudwig/echo-sample/app/driver/db/seed"
	"gorm.io/gorm"
)

type UsersRepository struct {
	db *gorm.DB
}

type User struct {
	IdField
	Name string
	TimeFields
}

func NewUsersRepository() (UsersRepository, error) {
	db, err := GetDbInstance()
	return UsersRepository{db: db}, err
}

func (u UsersRepository) GetAll() ([]user.User, error) {
	return seed.GetTestUsers()
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
