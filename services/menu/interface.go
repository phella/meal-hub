package menu

import "Bete/models"

type Service interface {
	// CreateMeal() Meal
	GetMenu(tableID uint) (Menu, error)
}

type Meal struct {
	Title       string
	Description string
	ImageLinks  []string
	Rating      float32
}

type Menu struct {
	Items map[string][]Meal
}

type CombinedMeal struct {
	Menu models.Menu
	Meal models.Meal
}
