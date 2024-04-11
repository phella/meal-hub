package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Id      uint    `gorm:"primaryKey;autoIncrement"`
	Session Session `gorm:"index:order__session_id__user_id"`
	User    User    `gorm:"index:order__session_id__user_id"`
	Meal    []Meal
}
