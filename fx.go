package main

import (
	"Bete/models"
	orderRouter "Bete/routes/order"
	"Bete/routes/restaurant"
	"Bete/services/Restaurant"
	"Bete/services/database"
	"Bete/services/order"
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
		fx.Provide(order.New),
		fx.Provide(orderRouter.New),

		// Invoke initialization methods
		fx.Invoke(models.MigrateSchema),
		fx.Invoke(router.Service.ListenAndServe),
	)
}
