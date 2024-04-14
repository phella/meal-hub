package orderRouter

import (
	"Bete/pkg/httputils"
	"Bete/services/order"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"net/http"
	"net/url"
	"strconv"
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
	router.Post("/items", r.addItems)
	router.Get("{order-id}/full-check", r.getFullCheck)
	router.Get("{order-id}/personal-check/{user-id}", r.getPersonalCheck)
	router.Get("{order-id}/split-check", r.getSplitCheck)
	router.Get("{order-id}/selected-check", r.getSelectedCheck)
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

	err = httputils.JSON(w, ordr)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (r orderRouter) getFullCheck(w http.ResponseWriter, req *http.Request) {
	param := chi.URLParam(req, "order-id")
	orderID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		http.Error(w, "invalid restaurant id", http.StatusBadRequest)
		return
	}

	check, err := r.orderService.CalculateFullCheck(req.Context(), order.CalculateFullCheckParams{
		OrderID: uint(orderID),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = httputils.JSON(w, struct {
		Check int64 `json:"check"`
	}{
		Check: check,
	}); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (r orderRouter) getPersonalCheck(w http.ResponseWriter, req *http.Request) {
	param := chi.URLParam(req, "order-id")
	orderID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		http.Error(w, "invalid restaurant id", http.StatusBadRequest)
		return
	}

	param = chi.URLParam(req, "user-id")
	userID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		http.Error(w, "invalid restaurant id", http.StatusBadRequest)
		return
	}

	check, err := r.orderService.CalculateUserCheck(req.Context(), order.CalculateUserCheckParams{
		OrderID: uint(orderID),
		UserID:  uint(userID),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = httputils.JSON(w, struct {
		Check int64 `json:"check"`
	}{
		Check: check,
	}); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (r orderRouter) getSplitCheck(w http.ResponseWriter, req *http.Request) {
	param := chi.URLParam(req, "order-id")
	orderID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		http.Error(w, "invalid restaurant id", http.StatusBadRequest)
		return
	}
	queryParam := req.URL.Query().Get("split-count")
	splitCount, err := strconv.ParseInt(queryParam, 10, 64)
	if err != nil {
		http.Error(w, "invalid restaurant id", http.StatusBadRequest)
		return
	}

	check, err := r.orderService.CalculateEquallyDividedCheck(req.Context(), order.CalculateEquallyDividedCheckParams{
		OrderID:     uint(orderID),
		SplitsCount: splitCount,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = httputils.JSON(w, struct {
		Check int64 `json:"check"`
	}{
		Check: check,
	}); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (r orderRouter) getSelectedCheck(w http.ResponseWriter, req *http.Request) {
	param := chi.URLParam(req, "order-id")
	orderID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		http.Error(w, "invalid restaurant id", http.StatusBadRequest)
		return
	}
	queryParam := req.URL.Query().Get("selected-items")
	jsonString, err := url.QueryUnescape(queryParam)
	if err != nil {
		http.Error(w, "failed to decode selected items", http.StatusBadRequest)
		return
	}

	var selectedItems []order.SelectedItems
	if err := json.Unmarshal([]byte(jsonString), &selectedItems); err != nil {
		http.Error(w, "failed to parse selected items", http.StatusBadRequest)
		return
	}

	check, err := r.orderService.CalculateSpecificCheckItems(req.Context(), order.CalculateSpecificCheckItemsParams{
		OrderID:       uint(orderID),
		SelectedItems: selectedItems,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = httputils.JSON(w, struct {
		Check int64 `json:"check"`
	}{
		Check: check,
	}); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
