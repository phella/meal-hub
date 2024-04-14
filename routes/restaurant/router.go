package restaurantRouter

import (
	"Bete/pkg/httputils"
	restaurantService "Bete/services/restaurant"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
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
	router.Post("/menu-qrcode/{id}", r.updateQrCodeMenu)
	router.Post("/", r.createRestaurant)
	return router
}

func (r restaurantRouter) updateQrCodeMenu(w http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	imgLink := req.FormValue("img_link")
	res := r.restaurantService.UpdateQrCodeMenu(imgLink, id)

	err := httputils.JSON(w, res)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (r restaurantRouter) getRestaurant(w http.ResponseWriter, req *http.Request) {
	param := chi.URLParam(req, "id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		http.Error(w, "invalid restaurant id", http.StatusBadRequest)
		return
	}
	res := r.restaurantService.GetRestaurant(id)

	err = httputils.JSON(w, res)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
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

	err = httputils.JSON(w, id)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
