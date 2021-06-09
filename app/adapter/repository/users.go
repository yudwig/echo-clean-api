package repository

import (
	"github.com/yudwig/echo-sample/app/core/domain/entity/user"
)

type Users interface {
	Get(id string) (user.User, error)
	GetAll() ([]user.User, error)
	Create(name string) (user.User, error)
	Delete(id string) error
	Update(id string, name string) (user.User, error)
}
