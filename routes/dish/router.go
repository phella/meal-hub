package dishRouter

import (
	"Bete/pkg/httputils"
	"Bete/services/Dish"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"net/http"
	"strconv"
)

type params struct {
	fx.In
	DishService dishService.Service
}

type Router interface {
	RegisterRoutes() chi.Router
}

type dishRouter struct {
	dishService dishService.Service
}

func New(p params) Router {
	return &dishRouter{
		dishService: p.DishService,
	}
}

func (d dishRouter) RegisterRoutes() chi.Router {
	router := chi.NewRouter()
	router.Get("/{id}", d.getDish)
	router.Post("/", d.createDish)
	return router
}

func (r dishRouter) getDish(w http.ResponseWriter, req *http.Request) {
	param := chi.URLParam(req, "id")
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		http.Error(w, "invalid dish id", http.StatusBadRequest)
		return
	}
	res := r.dishService.GetDish(id)

	_ = httputils.JSON(w, res)
}

func (r dishRouter) createDish(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")

	id := r.dishService.CreateDish(dishService.CreateDishParams{
		Name: name,
	})

	httputils.JSON(w, id)
}
