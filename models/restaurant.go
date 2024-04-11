package models

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	Id       uint `gorm:"primaryKey;autoIncrement"`
	Name     string
	LogoPath string
	Slogan   string
}
