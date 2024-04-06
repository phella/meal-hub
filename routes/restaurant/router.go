package restaurantRouter

import (
	"Bete/services/Restaurant"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"net/http"
	"strconv"
)

type params struct {
	fx.In

	RestaurantService restaurantService.Service
}

type Router interface {
	RegisterRoutes() chi.Router
}

type restaurantRouter struct {
	restaurantService restaurantService.Service
}

func New(p params) Router {
	return &restaurantRouter{
		restaurantService: p.RestaurantService,
	}
}

func (r restaurantRouter) RegisterRoutes() chi.Router {
	router := chi.NewRouter()
	router.Get("/{id}", r.getRestaurant)
	return router
}

func (r restaurantRouter) getRestaurant(w http.ResponseWriter, req *http.Request) {
	param := chi.URLParam(req, "id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		http.Error(w, "invalid restaurant id", http.StatusBadRequest)
		return
	}
	res := r.restaurantService.GetRestaurant(id)
	jsonData, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "error marshaling json", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
