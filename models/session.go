package models

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	Id       uint  `gorm:"primaryKey;autoIncrement"`
	IsActive *bool `gorm:"index:session__is_active"`
	Table    Table
}
