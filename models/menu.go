package models

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	ID           uint `gorm:"primaryKey;autoIncrement"`
	RestaurantID uint
	Meals        []Meal `gorm:"many2many:menu_meals;"`
	IsActive     *bool  `gorm:"uniqueIndex:"`
}
