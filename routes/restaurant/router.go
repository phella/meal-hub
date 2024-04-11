package restaurantRouter

import (
	"Bete/pkg/httputils"
	"Bete/services/restaurant"
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
	router.Post("/", r.createRestaurant)
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

	_ = httputils.JSON(w, res)
}

func (r restaurantRouter) createRestaurant(w http.ResponseWriter, req *http.Request) {
	err := req.ParseMultipartForm(30 << 20) // 30 MB max file size
	if err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		return
	}

	file, _, err := req.FormFile("logo")
	if err != nil {
		http.Error(w, "Unable to get logo from data", http.StatusBadRequest)
		return
	}
	defer file.Close()

	name := req.FormValue("name")
	slogan := req.FormValue("slogan")

	id := r.restaurantService.CreateRestaurant(restaurantService.CreateRestaurantParams{
		Name:   name,
		Slogan: slogan,
		Logo:   &file,
	})

	httputils.JSON(w, id)
}
