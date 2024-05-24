package menu

import (
	"Bete/services/database"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type menuService struct {
	db *gorm.DB
}

type params struct {
	fx.In
	DbService database.Service
}

func New(p params) Service {
	return &menuService{
		db: p.DbService.GetDBInstance(),
	}
}

func (s menuService) GetMenu(tableID string) (Menu, error) {
	var meals []CombinedMeal
	s.db.Raw(_getRestaurantMenu, tableID).Scan(&meals)
	return toMenu(meals)
}

func toMenu(meals []CombinedMeal) (Menu, error) {
	items := make(map[string][]Meal, len(meals))
	for _, meal := range meals {
		item := Meal{
			Title:       meal.Meal.Name,
			Description: meal.Meal.Desc,
		}
		items[meal.Meal.Tag] = append(items[meal.Meal.Tag], item)
	}
	return Menu{
		Items: items,
	}, nil
}
