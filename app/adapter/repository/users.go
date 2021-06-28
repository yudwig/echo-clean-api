package repository

import (
	"github.com/yudwig/echo-clean-api/app/core/domain/entity/user"
)

type Users interface {
	GetAll() ([]user.User, error)
	Create(name string) (user.User, error)
	Delete(id string) error
	Update(id string, name string) (user.User, error)
	Transaction(f func() error) error
}
