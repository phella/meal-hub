package main

import (
	"Bete/models"
	"Bete/routes/dish"
	"Bete/routes/restaurant"
	"Bete/services/Dish"
	"Bete/services/Restaurant"
	"Bete/services/database"
	"Bete/services/router"
	"go.uber.org/fx"
)

func InitializeApp() *fx.App {
	return fx.New(
		// Inject all the used services and routers
		fx.Provide(database.New),
		fx.Provide(router.New),
		fx.Provide(restaurantService.New),
		fx.Provide(restaurantRouter.New),
		fx.Provide(dishService.New),
		fx.Provide(dishRouter.New),

		// Invoke initialization methods
		fx.Invoke(models.MigrateSchema),
		fx.Invoke(router.Service.ListenAndServe),
	)
}
