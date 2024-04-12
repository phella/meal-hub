package order

import (
	"Bete/models"
	"Bete/pkg/pointers"
	"Bete/services/database"
	"context"
	"errors"
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
	result := s.db.Where("orders.is_active = ? AND orders.table_id = ?", true, tableId).First(&order)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Order{}, result.Error
		}
		newOrder := models.Order{
			IsActive: pointers.Ptr(true),
			TableId:  tableId,
		}
		res := s.db.Create(&newOrder)
		if res.Error != nil {
			return models.Order{}, res.Error
		}
		return newOrder, nil
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
	fmt.Println(meals[0].ID)
	fmt.Println(ID)
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

func toModelSelections(selections []Selection) []models.Selection {
	res := make([]models.Selection, len(selections))
	for i, selection := range selections {
		res[i] = models.Selection{
			ID: selection.ID,
		}
	}

	return res
}

func toOrderItems(orderItems []models.OrderItem) []OrderItem {
	res := make([]OrderItem, len(orderItems))
	for i, _ := range orderItems {
		res[i] = OrderItem{
			Meal:     orderItems[i].Meal,
			User:     orderItems[i].User,
			Quantity: orderItems[i].Quantity,
		}
	}

	return res
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
