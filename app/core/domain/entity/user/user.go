package user

import "encoding/json"

type User struct {
	Id   string
	Name string
}

func NewUser(Id string, Name string) *User {
	return &User{
		Id:   Id,
		Name: Name,
	}
}

func (u User) String() string {
	res, _ := json.Marshal(u)
	return string(res)
}
