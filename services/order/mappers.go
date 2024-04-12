package order

import "Bete/models"

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
