package rdb

func Migrate() {
	db, err := GetDbInstance()
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(UserModel{})
}
