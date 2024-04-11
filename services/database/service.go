package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type databaseService struct {
	db *gorm.DB
}

func New() Service {
	db, err := gorm.Open(mysql.Open("root@tcp(localhost:3306)/meal_hub?parseTime=true"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	return &databaseService{
		db: db,
	}
}

func (s databaseService) GetDBInstance() *gorm.DB {
	return s.db
}
