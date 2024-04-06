package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type databaseService struct {
	db *gorm.DB
}

func New() Service {
	db, err := gorm.Open(sqlite.Open("foo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &databaseService{
		db: db,
	}
}

func (s databaseService) GetDBInstance() *gorm.DB {
	return s.db
}
