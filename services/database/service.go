package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type databaseService struct {
	db *gorm.DB
}

func getDBConnectionString() string {
	godotenv.Load()
	db_name := os.Getenv("DB_NAME")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_hostname := os.Getenv("DB_HOSTNAME")
	db_port := os.Getenv("DB_PORT")
	db_portocol := os.Getenv("DB_PROTOCOL")
	return fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true", db_user, db_password, db_portocol, db_hostname, db_port, db_name)
}
func New() Service {
	dbConnectionString := getDBConnectionString()
	db, err := gorm.Open(mysql.Open(dbConnectionString), &gorm.Config{})
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
