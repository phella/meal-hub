package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID          uint  `gorm:"primaryKey;autoIncrement"`
	State       State `gorm:"default:0"`
	IsActive    *bool `gorm:"index:,unique,composite:order__table_id__is_active_idx"`
	TableId     uint  `gorm:"index:,unique,composite:order__table_id__is_active_idx"`
	CheckSplits *int64
	OrderItems  []OrderItem
}

type State int

const (
	Placed State = iota
	PartiallyPaid
	FullyPaid
	PaymentSplit
)
