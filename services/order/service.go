package order

import (
	"context"
	"fmt"
	"go.uber.org/fx"
)

type params struct {
	fx.In
}

type orderService struct {
}

func New(p params) Service {
	return &orderService{}
}

func (s orderService) AddItem(context.Context, AddItemParams) (Order, error) {
	return Order{}, fmt.Errorf("unimplemented")
}

func (s orderService) CalculateUserCheck(context.Context) (int64, error) {
	return 0, fmt.Errorf("unimplemented")
}

func (s orderService) CalculateFullCheck(context.Context) (int64, error) {
	return 0, fmt.Errorf("unimplemented")
}
