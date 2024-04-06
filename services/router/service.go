package router

import (
	"Bete/routes/restaurant"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"net/http"
	"os"
)

type routerService struct {
	restaurantRouter restaurantRouter.Router
}

type params struct {
	fx.In

	RestaurantRouter restaurantRouter.Router
}

func New(p params) Service {
	return &routerService{
		restaurantRouter: p.RestaurantRouter,
	}
}

func (s routerService) registerRoutes() chi.Router {
	mainRouter := chi.NewRouter()
	restaurantRoutes := s.restaurantRouter.RegisterRoutes()
	mainRouter.Mount("/api/v1/restaurant", restaurantRoutes)

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
