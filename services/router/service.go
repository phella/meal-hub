package router

import (
	orderRouter "Bete/routes/order"
	"Bete/routes/restaurant"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"net/http"
	"os"
)

type routerService struct {
	restaurantRouter restaurantRouter.Router
	orderRouter      orderRouter.Router
}

type params struct {
	fx.In

	RestaurantRouter restaurantRouter.Router
	OrderRouter      orderRouter.Router
}

func New(p params) Service {
	return &routerService{
		restaurantRouter: p.RestaurantRouter,
		orderRouter:      p.OrderRouter,
	}
}

func (s routerService) registerRoutes() chi.Router {
	mainRouter := chi.NewRouter()
	restaurantRoutes := s.restaurantRouter.RegisterRoutes()
	mainRouter.Mount("/api/v1/restaurant", restaurantRoutes)

	orderRoutes := s.orderRouter.RegisterRoutes()
	mainRouter.Mount("/api/v1/order", orderRoutes)

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
