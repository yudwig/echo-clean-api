package usecase

type UserUseCase interface {
	CreateUser(name string) error
}
