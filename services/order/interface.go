package order

import (
	"Bete/models"
	"context"
)

type Service interface {
	AddItems(context.Context, AddItemParams) (Order, error)
	CalculateUserCheck(context.Context) (int64, error)
	CalculateFullCheck(context.Context) (int64, error)
}

type Order struct {
	Id         uint
	OrderItems []OrderItem
}

type OrderItem struct {
	Meal     models.Meal // TODO: Create DTO objects for meal and user
	User     models.User
	Quantity int64
}

type AddItemParams struct {
	TableID uint
	UserID  uint
	Meals   []Meal
}

type Meal struct {
	ID       uint   `json:"id"`
	Dishes   []Dish `json:"dishes"`
	Quantity int64  `json:"quantity"`
}

type Dish struct {
	ID uint `json:"id"`
}

type insertOrderItemsParams struct {
	OrderID uint
	UserID  uint
	Meals   []Meal
}
