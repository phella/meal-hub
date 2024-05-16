package models

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Name     string
	LogoPath string
	Slogan   string
}
