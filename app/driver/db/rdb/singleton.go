package rdb

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SingletonDbInstance struct {
	db  *gorm.DB
	err error
}

var singletonDbInstance = newSingletonDb()

func newSingletonDb() SingletonDbInstance {
	dsn := "root@tcp(127.0.0.1:3306)/echo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return SingletonDbInstance{
		db:  db,
		err: err,
	}
}

func GetDbInstance() (*gorm.DB, error) {
	return singletonDbInstance.db, singletonDbInstance.err
}
