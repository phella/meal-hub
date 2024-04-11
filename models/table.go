package models

import "gorm.io/gorm"

type Table struct {
	gorm.Model
	Id    uint `gorm:"primaryKey;autoIncrement"`
	Order []Order
}
