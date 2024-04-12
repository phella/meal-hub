package restaurantService

import (
	"Bete/models"
	"Bete/services/database"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"mime/multipart"
)

type restaurantService struct {
	db *gorm.DB
}

type params struct {
	fx.In

	DbService database.Service
}

func New(p params) Service {
	return &restaurantService{
		db: p.DbService.GetDBInstance(),
	}
}

func (s restaurantService) CreateRestaurant(p CreateRestaurantParams) Restaurant {
	restaurant := Restaurant{
		Name:     p.Name,
		LogoPath: SaveImage(p.Logo),
	}
	s.db.Create(restaurant)
	return restaurant
}

func (s restaurantService) GetRestaurant(int64) Restaurant {
	var res models.Restaurant
	s.db.First(&res, 1)

	return Restaurant{
		Id:       res.Id,
		Name:     res.Name,
		LogoPath: res.LogoPath,
		Slogan:   res.Slogan,
	}
}

func (s restaurantService) CreateBranch() {}
func (s restaurantService) GetBranches()  {}

// TODO(mazen): Add any saving image algorithm
func SaveImage(image *multipart.File) string {
	return "img"
}
