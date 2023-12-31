package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func DbinstanceHandler(dbURL string) (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open(dbURL), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		return nil, err
	}

	return db, nil

}
