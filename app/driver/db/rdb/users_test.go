package rdb

import (
	"errors"
	"github.com/yudwig/echo-sample/app/core/domain/entity/user"
	"github.com/yudwig/ermux"
	"strconv"
	"testing"
)

func TestCreate(test *testing.T) {
	test.Run("create 1 user", func(t *testing.T) {
		_ = repos.users.Transaction(func() error {
			usr, err := repos.users.Create("user 1")
			if err != nil {
				t.Error(err)
				t.Log(usr)
				return err
			}

			users, _ := repos.users.GetAll()
			if len(users) != 1 {
				err = errors.New("row count not equal 1")
				t.Error(err)
				t.Log(users)
				return err
			}

			if usr.Name != users[0].Name {
				err = errors.New("unmatched user name")
				t.Error(err)
				t.Log("usr.Name: ", usr.Name)
				t.Log("users[0].Name: ", users[0].Name)
				return err
			}
			return errors.New("rollback")
		})
	})

	test.Run("create 3 user", func(t *testing.T) {
		_ = repos.users.Transaction(func() error {
			errs := make([]error, 3)
			users := make([]user.User, 3)
			users[0], errs[0] = repos.users.Create("user 1")
			users[1], errs[1] = repos.users.Create("user 2")
			users[2], errs[2] = repos.users.Create("user 3")
			if ermux.Some(errs) {
				t.Error(ermux.First(errs))
				t.Log(users)
				return ermux.First(errs)
			}

			res, _ := repos.users.GetAll()
			if len(res) != 3 {
				err := errors.New("row count not equal 3")
				t.Error(err)
				t.Log(res)
				return err
			}

			for i := range users {
				if users[i].Name != res[i].Name {
					err := errors.New("unmatched user name index: " + strconv.Itoa(i))
					t.Error(err)
					t.Log("usr.Name: ", users[i].Name)
					t.Log("res.Name: ", res[i].Name)
					return err
				}
			}
			return errors.New("rollback")
		})
	})

}

func TestDelete(test *testing.T) {
	test.Run("create -> delete", func(t *testing.T) {
		_ = repos.users.Transaction(func() error {
			// register 3 users
			errs := make([]error, 3)
			users := make([]user.User, 3)
			users[0], errs[0] = repos.users.Create("user 1")
			users[1], errs[1] = repos.users.Create("user 2")
			users[2], errs[2] = repos.users.Create("user 3")
			if ermux.Some(errs) {
				t.Error(ermux.First(errs))
				t.Log(users)
				return ermux.First(errs)
			}
			// fetch all users
			res, _ := repos.users.GetAll()
			if len(res) != 3 {
				err := errors.New("row count not equal 3")
				t.Error(err)
				t.Log(users)
				return err
			}
			// delete user no.2
			err := repos.users.Delete(res[1].Id.Value())
			if err != nil {
				t.Error(err)
				t.Log("res[1]: ", res[1])
				return err
			}
			// fetch all users
			res, _ = repos.users.GetAll()
			if len(res) != 2 {
				err := errors.New("row count not equal 2")
				t.Error(err)
				t.Log(users)
				return err
			}
			// check stored user no.1 and no.3
			if res[0].Name.Value() != "user 1" || res[1].Name.Value() != "user 3" {
				err := errors.New("delete incorrect recode")
				t.Error(err)
				t.Log("res[0]: ", res[0])
				t.Log("res[1]: ", res[1])
				return err
			}
			return errors.New("rollback")
		})
	})
}

func TestUpdate(test *testing.T) {
	test.Run("create -> delete", func(t *testing.T) {
		_ = repos.users.Transaction(func() error {
			// register 3 users
			errs := make([]error, 3)
			users := make([]user.User, 3)
			users[0], errs[0] = repos.users.Create("user 1")
			users[1], errs[1] = repos.users.Create("user 2")
			users[2], errs[2] = repos.users.Create("user 3")
			if ermux.Some(errs) {
				t.Error(ermux.First(errs))
				t.Log(users)
				return ermux.First(errs)
			}
			// fetch all users
			res, _ := repos.users.GetAll()
			if len(res) != 3 {
				err := errors.New("row count not equal 3")
				t.Error(err)
				t.Log(users)
				return err
			}
			// update user no.1
			_, err := repos.users.Update(res[1].Id.Value(), "updated user 2")
			if err != nil {
				t.Error(err)
				t.Log("res[1]: ", res[1])
				return err
			}
			// fetch all users
			res, _ = repos.users.GetAll()
			if len(res) != 3 {
				err := errors.New("row count not equal 3")
				t.Error(err)
				t.Log(users)
				return err
			}
			// check user no.2's name
			if res[1].Name.Value() != "updated user 2" {
				err := errors.New("not updated name")
				t.Error(err)
				t.Log("res[1]: ", res[1])
				return err
			}
			return errors.New("rollback")
		})
	})
}
