package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID         uint  `gorm:"primaryKey;autoIncrement"`
	IsActive   *bool `gorm:"index:,unique,composite:order__table_id__is_active_idx"`
	TableId    uint  `gorm:"index:,unique,composite:order__table_id__is_active_idx"`
	OrderItems []OrderItem
}
