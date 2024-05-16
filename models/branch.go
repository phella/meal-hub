package models

import "gorm.io/gorm"

type Branch struct {
	gorm.Model
	ID           uint `gorm:"primaryKey;autoIncrement"`
	RestaurantID uint
}
