package mockDataGenerator

import (
	"Bete/services/database"
	"go.uber.org/fx"
)

type params struct {
	fx.In
	DbService database.Service
}

func GenerateData(p params) {
	db := p.DbService.GetDBInstance()

	db.FirstOrCreate(&restaurant, restaurant)
	branch.RestaurantID = restaurant.ID
	db.FirstOrCreate(&branch, branch)
	menu.RestaurantID = restaurant.ID
	db.FirstOrCreate(&menu, menu)
}
