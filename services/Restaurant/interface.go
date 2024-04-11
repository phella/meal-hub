package restaurantService

import (
	"mime/multipart"
)

type Service interface {
	CreateRestaurant(CreateRestaurantParams) Restaurant
	GetRestaurant(int64) Restaurant
	CreateBranch()
	GetBranches()
}

type Restaurant struct {
	// TO ADD:
	// id       uint
	Name     string
	LogoPath string
	// TO ADD:
	// slogan   string
}

type CreateRestaurantParams struct {
	Name   string
	Logo   *multipart.File
	Slogan string
}
