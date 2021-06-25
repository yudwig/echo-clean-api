package rdb

import (
	"github.com/yudwig/ermux"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	errs := make([]error, 1)
	errs[0] = db.AutoMigrate(User{})
	if ermux.Some(errs) {
		panic(ermux.First(errs).Error())
	}
}

func MigrateWithDelete(db *gorm.DB) {
	errs := make([]error, 2)
	errs[0] = db.Migrator().DropTable(User{})
	errs[1] = db.AutoMigrate(User{})
	if ermux.Some(errs) {
		panic(ermux.First(errs).Error())
	}
}
