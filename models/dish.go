package models

import "gorm.io/gorm"

type Dish struct {
	gorm.Model
	Id   uint `gorm:"primaryKey;autoIncrement"`
	Name string
}
