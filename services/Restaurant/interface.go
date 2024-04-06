package restaurantService

import "gorm.io/gorm"

type Service interface {
	CreateRestaurant(CreateRestaurantParams) Restaurant
	GetRestaurant(int64) Restaurant
	CreateBranch()
	GetBranches()
}

type Restaurant struct {
	gorm.Model
	id       uint `gorm:"primaryKey;autoIncrement"`
	Name     string
	LogoPath string
}

type CreateRestaurantParams struct {
	Name     string
	LogoPath string
}
