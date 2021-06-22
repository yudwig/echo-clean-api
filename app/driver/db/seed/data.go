package seed

import (
	"github.com/yudwig/echo-sample/app/core/domain/entity/user"
	"github.com/yudwig/ermux"
)

func GetTestUsers() ([]user.User, error) {

	errs := make([]error, 3)
	users := make([]user.User, 3)

	users[0], errs[0] = user.NewUser("id-1", "user 1")
	users[1], errs[1] = user.NewUser("id-2", "user 2")
	users[2], errs[2] = user.NewUser("id-3", "user 3")

	if ermux.Some(errs) {
		return []user.User{}, ermux.First(errs)
	}

	return []user.User{
		users[0],
		users[1],
		users[2],
	}, nil
}
