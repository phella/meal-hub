package order

import (
	"Bete/models"
	"Bete/pkg/pointers"
	"Bete/services/database"
	"context"
	"errors"
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

func (s orderService) CalculateUserCheck(ctx context.Context, p CalculateUserCheckParams) (int64, error) {
	var orderItems []models.OrderItem
	res := s.db.Where(models.OrderItem{UserID: p.UserID, OrderID: p.OrderID}).Where("quantity != paid_quantity").Find(&orderItems)
	if res.Error != nil {
		return 0, res.Error
	}
	return calculateCheck(orderItems, true), nil
}

func (s orderService) CalculateFullCheck(ctx context.Context, p CalculateFullCheckParams) (int64, error) {
	var orderItems []models.OrderItem
	res := s.db.Where(models.OrderItem{OrderID: p.OrderID}).Where("quantity != paid_quantity").Find(&orderItems)
	if res.Error != nil {
		return 0, res.Error
	}
	return calculateCheck(orderItems, true), nil
}

func (s orderService) CalculateSpecificCheckItems(ctx context.Context, params CalculateSpecificCheckItemsParams) (int64, error) {
	var orderItems []models.OrderItem
	res := s.db.Where("ID IN ? AND order_id = ? = ", getOrderItemsIDs(params), params.OrderID).Find(&orderItems)
	if res.Error != nil {
		return 0, res.Error
	}

	if !ValidateCheckItems(orderItems, params) {
		return 0, errors.New("invalid items selection")
	}

	return calculateCheck(orderItems, true), nil
}

func (s orderService) CalculateEquallyDividedCheck(ctx context.Context, p CalculateEquallyDividedCheckParams) (int64, error) {
	var orderItems []models.OrderItem
	res := s.db.Where(models.OrderItem{OrderID: p.OrderID}).Find(&orderItems)
	if res.Error != nil {
		return 0, res.Error
	}

	// TODO(Philo): use money package
	return calculateCheck(orderItems, false) / p.SplitsCount, nil
}

func calculateCheck(items []models.OrderItem, removePaidItems bool) int64 {
	check := int64(0)
	for _, item := range items {
		quantity := item.Quantity
		if removePaidItems {
			quantity -= item.PaidQuantity
		}
		check += item.PriceE5 * quantity
	}

	return check
}

func (s orderService) ensureOrder(tableId uint) (models.Order, error) {
	order := models.Order{
		IsActive: pointers.Ptr(true),
		TableId:  tableId,
	}
	result := s.db.Where(order).FirstOrCreate(&order)
	if result.Error != nil {
		return models.Order{}, result.Error
	}

	return order, nil
}

func (s orderService) insertOrderItems(params insertOrderItemsParams) error {

	var selectionsInfo []models.Selection
	s.db.Where("ID IN ?", getSelectionsIDs(params.Meals)).Find(&selectionsInfo)

	var mealsInfo []models.Meal
	s.db.Where("ID IN ?", getMealsIDs(params.Meals)).Find(&mealsInfo)

	items := make([]models.OrderItem, len(params.Meals))
	for i, meal := range params.Meals {

		priceE5, err := s.calculateMealPriceE5(meal, mealsInfo, selectionsInfo)
		if err != nil {
			return err
		}

		items[i] = models.OrderItem{
			OrderID:    params.OrderID,
			UserID:     params.UserID,
			MealID:     meal.ID,
			PriceE5:    priceE5,
			Quantity:   meal.Quantity,
			Selections: toModelSelections(meal.Selections),
		}
	}

	res := s.db.Create(&items)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (s orderService) calculateMealPriceE5(meal Meal, mealsInfo []models.Meal, selectionsInfo []models.Selection) (int64, error) {
	price, err := findMealPriceE5ByID(mealsInfo, meal.ID)
	if err != nil {
		return 0, err
	}

	for _, selection := range selectionsInfo {
		extraCharge, err := findSelectionExtraChargeE5ByID(selectionsInfo, selection.ID)
		if err != nil {
			return 0, err
		}
		price += extraCharge
	}
	return price, nil
}

func findMealPriceE5ByID(meals []models.Meal, ID uint) (int64, error) {
	for _, meal := range meals {
		if meal.ID == ID {
			return meal.PriceE5, nil
		}
	}

	return 0, errors.New("meal not found")
}

func findSelectionExtraChargeE5ByID(selections []models.Selection, ID uint) (int64, error) {
	for _, selection := range selections {
		if selection.ID == ID {
			return selection.ExtraChargesE5, nil
		}
	}

	return 0, errors.New("selection not found")
}

func getMealsIDs(meals []Meal) []uint {
	mealIDs := make([]uint, len(meals))
	for i, meal := range meals {
		mealIDs[i] = meal.ID
	}

	return mealIDs
}

func getSelectionsIDs(meals []Meal) []uint {
	selectionIDs := []uint{}
	for _, meal := range meals {
		for _, selection := range meal.Selections {
			selectionIDs = append(selectionIDs, selection.ID)
		}
	}

	return selectionIDs
}

func getOrderItemsIDs(p CalculateSpecificCheckItemsParams) []uint {
	orderItemIds := []uint{}
	for _, orderItem := range p.SelectedItems {
		orderItemIds = append(orderItemIds, orderItem.ID)
	}

	return orderItemIds
}

func ValidateCheckItems(orderItems []models.OrderItem, params CalculateSpecificCheckItemsParams) bool {

	if len(orderItems) != len(params.SelectedItems) {
		return false
	}

	for _, orderItem := range orderItems {
		for _, selectedOrderItem := range params.SelectedItems {
			unpaidQuantity := orderItem.Quantity - orderItem.PaidQuantity
			if orderItem.ID == selectedOrderItem.ID && selectedOrderItem.Quantity > unpaidQuantity {
				return false
			}
		}
	}

	return true
}

func (s orderService) getHydratedOrder(orderID uint) (Order, error) {
	var orderItems []models.OrderItem
	res := s.db.Joins("Meal").Joins("User").Where(models.OrderItem{OrderID: orderID}).Find(&orderItems)
	if res.Error != nil {
		return Order{}, res.Error
	}

	return Order{
		Id:         orderID,
		OrderItems: toOrderItems(orderItems),
	}, nil
}
