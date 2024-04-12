package orderRouter

import (
	"Bete/pkg/httputils"
	"Bete/services/order"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"net/http"
)

type params struct {
	fx.In

	OrderService order.Service
}

type Router interface {
	RegisterRoutes() chi.Router
}

type orderRouter struct {
	orderService order.Service
}

func New(p params) Router {
	return &orderRouter{
		orderService: p.OrderService,
	}
}

func (r orderRouter) RegisterRoutes() chi.Router {
	router := chi.NewRouter()
	// router.Get("/{id}", r.getRestaurant)
	router.Post("/items", r.addItems)
	return router
}

func (r orderRouter) addItems(w http.ResponseWriter, req *http.Request) {
	type AddItemReq struct {
		TableID uint         `json:"table_id"`
		Meals   []order.Meal `json:"meals"`
	}
	var addItemReq AddItemReq
	err := json.NewDecoder(req.Body).Decode(&addItemReq)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	ordr, err := r.orderService.AddItems(req.Context(), order.AddItemParams{
		TableID: addItemReq.TableID,
		UserID:  1,
		Meals:   addItemReq.Meals,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	httputils.JSON(w, ordr)
}
