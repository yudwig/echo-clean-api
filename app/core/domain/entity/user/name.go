package user

import "errors"

type Name struct {
	value string
}

func NewName(name string) (Name, error) {
	if len(name) == 0 {
		return Name{}, errors.New(EmptyUserNameError)
	}
	return Name{
		value: name,
	}, nil
}

func (n Name) Get() string {
	return n.value
}
