package dishService

type Service interface {
	CreateDish(CreateDishParams) Dish

	GetDish(int64) Dish
}
type Dish struct {
	id   uint
	Name string
}

type CreateDishParams struct {
	Name string
}
