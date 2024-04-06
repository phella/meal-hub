package account

type Service interface {
	CreateMeal() Meal
	
}

type Meal struct {
	Title       string
	Description string
	ImageLinks  []string
	Rating      float32
}
