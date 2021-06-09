package user

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
