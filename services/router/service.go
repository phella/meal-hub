package router

import (
	"Bete/routes/restaurant"
	"Bete/routes/user"
	"fmt"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"net/http"
	"os"
)

type routerService struct {
	restaurantRouter restaurantRouter.Router
	userRouter       userRouter.Router
}

type params struct {
	fx.In

	RestaurantRouter restaurantRouter.Router
	UserRouter       userRouter.Router
}

func New(p params) Service {
	return &routerService{
		restaurantRouter: p.RestaurantRouter,
		userRouter:       p.UserRouter,
	}
}

func (s routerService) registerRoutes() chi.Router {
	mainRouter := chi.NewRouter()
	restaurantRoutes := s.restaurantRouter.RegisterRoutes()
	mainRouter.Mount("/api/v1/restaurant", restaurantRoutes)

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
