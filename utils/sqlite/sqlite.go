package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewConnection(file string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
