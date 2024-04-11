package models

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	id       uint `gorm:"primaryKey;autoIncrement"`
	Name     string
	LogoPath string
	slogan   string
}
