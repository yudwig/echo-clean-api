package user

import (
	"encoding/json"
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

func (i Id) Value() string {
	return i.value
}

func (i Id) String() string {
	return i.value
}

func (i Id) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Value())
}
