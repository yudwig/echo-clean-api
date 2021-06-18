package rdb

import (
	"gorm.io/gorm"
	"time"
)

type IdField struct {
	Id uint `gorm:"primarykey"`
}

type TimeFields struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
