package usecase

type UserUseCase interface {
	CreateUser(name string) error
	GetUsers() error
	GetUser(id string) error
	DeleteUser(id string) error
	UpdateUserName(id string, name string) error
}
