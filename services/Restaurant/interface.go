package restaurantService

type Service interface {
	CreateRestaurant(CreateRestaurantParams) Restaurant
	GetRestaurant(int64) Restaurant
}

type Restaurant struct {
	AddedID  int64  `db:"added_id"`
	Name     string `db:"name"`
	LogoPath string `db:"logo_path"`
}

type CreateRestaurantParams struct {
	Name     string
	LogoPath string
}
