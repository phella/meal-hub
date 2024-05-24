package models

import "gorm.io/gorm"

type Meal struct {
	gorm.Model
	ID             uint `gorm:"primaryKey;autoIncrement"`
	Name           string
	Desc           string
	ImagePath      string
	Tag            string // TODO(Philo): create a tag table to support more than 1 tag.
	PriceE5        int64
	IsCustomizable bool
}
