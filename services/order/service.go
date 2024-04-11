package order

import (
	"Bete/models"
	"Bete/pkg/pointers"
	"Bete/services/database"
	"context"
	"fmt"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type params struct {
	fx.In

	DbService database.Service
}

type orderService struct {
	db *gorm.DB
}

func New(p params) Service {
	return &orderService{
		db: p.DbService.GetDBInstance(),
	}
}

func (s orderService) AddItems(ctx context.Context, params AddItemParams) (Order, error) {
	order, err := s.ensureOrder(params.TableID)
	if err != nil {
		return Order{}, err
	}

	if err := s.insertOrderItems(insertOrderItemsParams{
		OrderID: order.ID,
		UserID:  params.UserID,
		Meals:   params.Meals,
	}); err != nil {
		return Order{}, err
	}

	return s.getHydratedOrder(order.ID)
}

func (s orderService) CalculateUserCheck(context.Context) (int64, error) {
	return 0, fmt.Errorf("unimplemented")
}

func (s orderService) CalculateFullCheck(context.Context) (int64, error) {
	return 0, fmt.Errorf("unimplemented")
}

func (s orderService) ensureOrder(tableId uint) (models.Order, error) {
	var order models.Order
	result := s.db.Where("is_active = ? AND table_id = ?", true, tableId).First(&order)
	if result.Error != nil {
		newOrder := models.Order{
			IsActive: pointers.Ptr(true),
			TableId:  tableId,
		}
		res := s.db.Create(newOrder)
		if res.Error != nil {
			return models.Order{}, res.Error
		}
		return newOrder, nil
	}

	return order, nil
}

func (s orderService) insertOrderItem(item models.OrderItem) error {
	res := s.db.Create(item)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (s orderService) insertOrderItems(params insertOrderItemsParams) error {
	for _, meal := range params.Meals {
		priceE5, err := s.getMealPriceE5(meal)
		if err != nil {
			return err
		}

		item := models.OrderItem{
			OrderID:  params.OrderID,
			UserID:   params.UserID,
			MealID:   meal.ID,
			PriceE5:  priceE5,
			Quantity: meal.Quantity,
			Dishes:   toModelDishes(meal.Dishes),
		}
		if err := s.insertOrderItem(item); err != nil {
			return err
		}
	}

	return nil
}

func (s orderService) getMealPriceE5(meal Meal) (int64, error) {
	return 33, nil
}

func toModelDishes(dishes []Dish) []models.Dish {
	res := make([]models.Dish, len(dishes))
	for i, dish := range dishes {
		res[i] = models.Dish{
			ID: dish.ID,
		}
	}

	return res
}

func toOrderItems(orderItems []models.OrderItem) []OrderItem {
	res := make([]OrderItem, len(orderItems))
	for i, orderItem := range orderItems {
		res[i] = OrderItem{
			Meal:     orderItem.Meal,
			User:     orderItem.User,
			Quantity: orderItem.Quantity,
		}
	}

	return res
}

func (s orderService) getHydratedOrder(orderID uint) (Order, error) {
	var order models.Order
	res := s.db.First(&order, orderID)
	if res.Error != nil {
		return Order{}, res.Error
	}

	return Order{
		Id:         orderID,
		OrderItems: toOrderItems(order.OrderItems),
	}, nil
}
