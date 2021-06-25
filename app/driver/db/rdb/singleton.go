package rdb

import (
	"flag"
	"gorm.io/gorm"
)

type SingletonDbInstance struct {
	db  *gorm.DB
	err error
}

var singletonDbInstance = newSingletonDb()

func newSingletonDb() SingletonDbInstance {

	var (
		db  *gorm.DB
		err error
	)
	if flag.Lookup("test.v") != nil {
		db, err = NewTestDb()
	}
	db, err = NewDb()

	return SingletonDbInstance{
		db:  db,
		err: err,
	}
}

func GetDbInstance() (*gorm.DB, error) {
	return singletonDbInstance.db, singletonDbInstance.err
}
