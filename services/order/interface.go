package order

import "context"

type Service interface {
	AddItem(context.Context, AddItemParams) (Order, error)
	CalculateUserCheck(context.Context) (int64, error)
	CalculateFullCheck(context.Context) (int64, error)
}

type Order struct {
	Id       int64
	PriceE5  int64
	quantity int64
}

type OrderItem struct {
}

type AddItemParams struct {
}
