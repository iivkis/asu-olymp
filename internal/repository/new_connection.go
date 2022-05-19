package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./lite.sqlite"))
	if err != nil {
		panic(err)
	}
	return db
}
