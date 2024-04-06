package restaurantService

import (
	"Bete/services/database"
	"go.uber.org/fx"
	"gorm.io/gorm"
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

func (s restaurantService) CreateRestaurant(CreateRestaurantParams) Restaurant {
	return Restaurant{}
}

func (s restaurantService) GetRestaurant(int64) Restaurant {
	var res Restaurant
	s.db.First(&res, 1)

	return res
}

func (s restaurantService) CreateBranch() {}
func (s restaurantService) GetBranches()  {}
