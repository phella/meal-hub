package main

import (
	"Bete/models"
	"Bete/routes/dish"
	menuRouter "Bete/routes/menu"
	orderRouter "Bete/routes/order"
	restaurantRouter "Bete/routes/restaurant"
	userRouter "Bete/routes/user"
	"Bete/services/Dish"
	"Bete/services/database"
	"Bete/services/menu"
	mockDataGenerator "Bete/services/mock_data_generator"
	"Bete/services/order"
	restaurantService "Bete/services/restaurant"
	"Bete/services/router"
	userService "Bete/services/user"

	"go.uber.org/fx"
)

func InitializeApp() *fx.App {
	return fx.New(
		// Inject all the used services and routers
		fx.Provide(database.New),
		fx.Provide(router.New),
		fx.Provide(restaurantService.New),
		fx.Provide(restaurantRouter.New),
		fx.Provide(menu.New),
		fx.Provide(menuRouter.New),
		fx.Provide(dishService.New),
		fx.Provide(dishRouter.New),
		fx.Provide(order.New),
		fx.Provide(orderRouter.New),
		fx.Provide(userService.New),
		fx.Provide(userRouter.New),

		// Invoke initialization methods
		fx.Invoke(models.MigrateSchema),
		fx.Invoke(mockDataGenerator.GenerateData),
		fx.Invoke(router.Service.ListenAndServe),
	)
}
