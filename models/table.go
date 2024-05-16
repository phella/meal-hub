package models

import "gorm.io/gorm"

type Table struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;autoIncrement"`
	BranchID uint
	Order    []Order
}
