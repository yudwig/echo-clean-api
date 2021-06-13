package user

import (
	"encoding/json"
)

type User struct {
	Id   Id
	Name Name
}

func NewUser(id string, name string) (User, error) {
	var (
		err      error
		userId   Id
		userName Name
	)
	userId, err = NewId(id)
	if err != nil {
		return User{}, err
	}
	userName, err = NewName(name)
	if err != nil {
		return User{}, err
	}
	return User{
		Id:   userId,
		Name: userName,
	}, nil
}

func (u User) String() string {
	res, _ := json.Marshal(u)
	return string(res)
}
