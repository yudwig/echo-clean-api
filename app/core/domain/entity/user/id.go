package user

import (
	"errors"
)

type Id struct {
	value string
}

func NewId(id string) (Id, error) {
	if len(id) == 0 {
		return Id{}, errors.New(InvalidUserIdError)
	}
	return Id{value: id}, nil
}

func (i Id) Get() string {
	return i.value
}
