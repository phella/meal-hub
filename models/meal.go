package models

import "gorm.io/gorm"

type Meal struct {
	gorm.Model
	ID             uint `gorm:"primaryKey;autoIncrement"`
	Name           string
	Desc           string
	ImagePath      string
	Tag            string
	PriceE5        int64
	IsCustomizable bool
}
