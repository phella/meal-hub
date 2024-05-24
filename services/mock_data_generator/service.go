package mockDataGenerator

import (
	"Bete/services/database"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type mockDataGenerator struct {
	db *gorm.DB
}

type params struct {
	fx.In
	DbService database.Service
}

func New(p params) Service {
	return &mockDataGenerator{
		db: p.DbService.GetDBInstance(),
	}
}

func (s mockDataGenerator) GenerateData() {
	s.db.FirstOrCreate(&restaurant, restaurant)
	branch.RestaurantID = restaurant.ID
	s.db.FirstOrCreate(&branch, branch)
	menu.RestaurantID = restaurant.ID
	s.db.FirstOrCreate(&menu, menu)
}
