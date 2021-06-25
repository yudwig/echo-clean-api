package rdb

import (
	"github.com/yudwig/echo-sample/app/core/domain/entity/user"
	"gorm.io/gorm"
	"strconv"
)

type UsersRepository struct {
	*Rdb
}

type User struct {
	IdField
	Name string
	TimeFields
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return UsersRepository{
		Rdb: &Rdb{db: db},
	}
}

func (u UsersRepository) GetAll() ([]user.User, error) {
	var users []User
	res := u.db.Find(&users)
	if res.Error != nil {
		return []user.User{}, res.Error
	}
	var entities []user.User
	for _, usr := range users {
		entity, err := user.NewUser(strconv.Itoa(int(usr.Id)), usr.Name)
		if err != nil {
			return []user.User{}, err
		}
		entities = append(entities, entity)
	}
	return entities, nil
}

func (u UsersRepository) Create(name string) (user.User, error) {
	usr := User{
		Name: name,
	}
	res := u.db.Create(&usr)
	if res.Error != nil {
		return user.User{}, res.Error
	}
	return user.NewUser(strconv.Itoa(int(usr.Id)), usr.Name)
}

func (u UsersRepository) Delete(id string) error {
	i, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	usr := User{
		IdField: IdField{
			Id: uint(i),
		},
	}
	u.db.Delete(&usr)
	return nil
}

func (u UsersRepository) Update(id string, name string) (user.User, error) {
	i, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return user.User{}, err
	}
	usr := User{
		IdField: IdField{
			Id: uint(i),
		},
	}
	u.db.Model(&usr).Update("name", name)
	return user.NewUser(id, name)
}
