package dishService

import (
	"Bete/models"
	"Bete/services/database"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type dishService struct {
	db *gorm.DB
}

type params struct {
	fx.In
	DbService database.Service
}

func New(p params) Service {
	return &dishService{
		db: p.DbService.GetDBInstance(),
	}
}

func (s dishService) CreateDish(dishParams CreateDishParams) Dish {
	dish := models.Dish{
		Name: dishParams.Name,
	}
	s.db.Create(&dish)
	return Dish{Name: dish.Name}
}

func (s dishService) GetDish(id int64) Dish {
	var dish Dish
	s.db.First(&dish, id)
	return dish
}
