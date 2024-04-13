package models

import "gorm.io/gorm"

type Selection struct {
	gorm.Model
	ID             uint `gorm:"primaryKey;autoIncrement"`
	SectionID      uint
	DishID         uint
	ExtraChargesE5 int64
}
