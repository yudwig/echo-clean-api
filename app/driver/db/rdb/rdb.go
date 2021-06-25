package rdb

import (
	"gorm.io/gorm"
)

type Rdb struct {
	db *gorm.DB
}

func (r *Rdb) Transaction(f func() error) error {

	return r.db.Transaction(func(tx *gorm.DB) error {
		// backup DB instance.
		db := r.db

		// temporarily replace DB with tx-DB instance.
		r.db = tx

		// exec callback
		err := f()

		// restore DB instance
		r.db = db
		return err
	})
}
