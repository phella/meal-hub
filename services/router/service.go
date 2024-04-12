package router

import (
	orderRouter "Bete/routes/order"
	restaurantRouter "Bete/routes/restaurant"
	userRouter "Bete/routes/user"
	userService "Bete/services/user"
	"context"
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
	userService      userService.Service
}

type params struct {
	fx.In

	RestaurantRouter restaurantRouter.Router
	OrderRouter      orderRouter.Router
	UserRouter       userRouter.Router
	UserService      userService.Service
}

func New(p params) Service {
	return &routerService{
		restaurantRouter: p.RestaurantRouter,
		orderRouter:      p.OrderRouter,
		userRouter:       p.UserRouter,
		userService:      p.UserService,
	}
}

func (s routerService) EnsureUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		type ContextKey string
		const ContextTokenKey ContextKey = "Id"
		token := r.Header.Get("Token")
		userParams := userService.EnsureUserParams{Token: token}
		user := s.userService.EnsureUser(userParams)
		ctx := context.WithValue(r.Context(), ContextTokenKey, user.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s routerService) registerRoutes() chi.Router {
	mainRouter := chi.NewRouter()
	restaurantRoutes := s.restaurantRouter.RegisterRoutes()
	mainRouter.Mount("/api/v1/restaurant", restaurantRoutes)

	orderRoutes := s.orderRouter.RegisterRoutes()
	mainRouter.With(s.EnsureUser).Mount("/api/v1/order", orderRoutes)

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
