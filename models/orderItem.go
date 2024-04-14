package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	Id           uint `gorm:"primaryKey;autoIncrement"`
	OrderID      uint
	UserID       uint
	MealID       uint
	PriceE5      int64
	Quantity     int64
	PaidQuantity int64       `gorm:"default:false"`
	Selections   []Selection `gorm:"many2many:item_selections;"`
	User         User        `gorm:"foreignKey:UserID"`
	Meal         Meal        `gorm:"foreignKey:MealID"`
}
