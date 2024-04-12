package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Token string `gorm:"uniqueIndex:,length:14"`
	Name  string
}
