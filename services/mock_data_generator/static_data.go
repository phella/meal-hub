package mockDataGenerator

import (
	"Bete/models"
	"Bete/pkg/pointers"
)

var (
	restaurant = models.Restaurant{
		Name:   "MoBistro",
		Slogan: "International, Grills, Italian",
	}

	branch = models.Branch{}

	menu = models.Menu{
		IsActive: pointers.Ptr(true),
		Meals: []models.Meal{
			{
				Name:           "Mo's Bolognese Spaghetti",
				Desc:           "Spaghetti, Slow-cooked minced beef, marinated in fresh basil and slow-cooked in Mo's signature tomato velvety sauce",
				Tag:            "Pasta",
				PriceE5:        int64(27500000),
				IsCustomizable: false,
			},
		},
	}
)
