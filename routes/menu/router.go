package menuRouter

import (
	"Bete/pkg/httputils"
	"Bete/services/menu"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

type params struct {
	fx.In

	MenuService menu.Service
}

type Router interface {
	RegisterRoutes() chi.Router
}

type menuRouter struct {
	menuService menu.Service
}

func New(p params) Router {
	return &menuRouter{
		menuService: p.MenuService,
	}
}

func (r menuRouter) RegisterRoutes() chi.Router {
	router := chi.NewRouter()
	router.Post("/tables/{table-id}", r.getMenu)
	return router
}

func (r menuRouter) getMenu(w http.ResponseWriter, req *http.Request) {
	tableID := chi.URLParam(req, "table-id")
	token, err := r.menuService.GetMenu(tableID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = httputils.JSON(w, token)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
