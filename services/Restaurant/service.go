package restaurantService

import (
	"Bete/services/database"
	"go.uber.org/fx"
)

type restaurantService struct {
	dbService database.Service
}

type params struct {
	fx.In

	DbService database.Service
}

func New(p params) Service {
	return &restaurantService{
		dbService: p.DbService,
	}
}

func (s restaurantService) CreateRestaurant(CreateRestaurantParams) Restaurant {
	return Restaurant{}
}

func (s restaurantService) GetRestaurant(int64) Restaurant {
	return Restaurant{
		Name: "Flafel Philo",
	}
}
