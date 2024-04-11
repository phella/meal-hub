package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID         uint  `gorm:"primaryKey;autoIncrement"`
	IsActive   *bool `gorm:"index:session__is_active"`
	TableId    uint
	OrderItems []OrderItem
}
