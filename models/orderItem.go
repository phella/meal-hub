package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	Id       uint `gorm:"primaryKey;autoIncrement"`
	User     User
	Order    Order
	Meal     Meal
	PriceE5  int64
	// TO ADD:
	//quantity int64
	Dish     Dish `gorm:"many2many:item_dishes;"`
}
