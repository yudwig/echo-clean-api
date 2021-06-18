package rdb

import "fmt"

func Migrate() {
	fmt.Println("Migrate start.")
	db, err := GetDbInstance()
	if err != nil {
		panic(err.Error())
	}
	err = db.AutoMigrate(User{})
	if err != nil {
		panic(err.Error())
	}
}
