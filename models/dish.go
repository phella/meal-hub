package models

import "gorm.io/gorm"

type Dish struct {
	gorm.Model
	ID   uint `gorm:"primaryKey;autoIncrement"`
	Name string
}
