package user

import (
	"encoding/json"
	"errors"
)

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

func (n Name) Value() string {
	return n.value
}

func (n Name) String() string {
	return n.value
}

func (n Name) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Value())
}
