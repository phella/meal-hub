package restaurantService

import (
	"mime/multipart"
)

type Service interface {
	CreateRestaurant(CreateRestaurantParams) Restaurant
	GetRestaurant(int64) Restaurant
	UpdateQrCodeMenu(string, string) string
	CreateBranch()
	GetBranches()
}

type Restaurant struct {
	Id       uint
	Name     string
	LogoPath string
	Slogan   string
}

type CreateRestaurantParams struct {
	Name   string
	Logo   *multipart.File
	Slogan string
}
