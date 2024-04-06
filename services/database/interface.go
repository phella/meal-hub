package database

import "gorm.io/gorm"

type Service interface {
	GetDBInstance() *gorm.DB
}
