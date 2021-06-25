package rdb

import "gorm.io/gorm"

type Repos struct {
	db    *gorm.DB
	users UsersRepository
}

var repos Repos

func init() {
	db, err := NewTestDb()
	if err != nil {
		panic(err)
	}
	repos = Repos{
		db:    db,
		users: NewUsersRepository(db),
	}
}
