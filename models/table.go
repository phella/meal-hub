package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Table struct {
	gorm.Model
	ID       string `gorm:"primaryKey;"`
	BranchID uint
	Order    []Order
}

func (table *Table) BeforeCreate(tx *gorm.DB) (err error) {
	table.ID = uuid.NewString()
	return
}
