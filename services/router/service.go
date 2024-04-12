package router

import (
	orderRouter "Bete/routes/order"
	restaurantRouter "Bete/routes/restaurant"
	userRouter "Bete/routes/user"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type routerService struct {
	restaurantRouter restaurantRouter.Router
	orderRouter      orderRouter.Router
	userRouter       userRouter.Router
}

type params struct {
	fx.In

	RestaurantRouter restaurantRouter.Router
	OrderRouter      orderRouter.Router
	UserRouter       userRouter.Router
}

func New(p params) Service {
	return &routerService{
		restaurantRouter: p.RestaurantRouter,
		orderRouter:      p.OrderRouter,
		userRouter:       p.UserRouter,
	}
}

func (s routerService) registerRoutes() chi.Router {
	mainRouter := chi.NewRouter()
	restaurantRoutes := s.restaurantRouter.RegisterRoutes()
	mainRouter.Mount("/api/v1/restaurant", restaurantRoutes)

	orderRoutes := s.orderRouter.RegisterRoutes()
	mainRouter.Mount("/api/v1/order", orderRoutes)

	userRoutes := s.userRouter.RegisterRoutes()
	mainRouter.Mount("/api/v1/user", userRoutes)

	return mainRouter
}

func (s routerService) ListenAndServe() {
	go func() {
		if err := http.ListenAndServe(":8080", s.registerRoutes()); err != nil {
			fmt.Println("failed to listen on port 8080")
			// TODO: test this
			os.Exit(0)
		}
	}()
}
