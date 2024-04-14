package restaurantService

import (
	"Bete/models"
	"Bete/services/database"
	"fmt"
	"mime/multipart"

	"github.com/skip2/go-qrcode"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type restaurantService struct {
	db *gorm.DB
}

type params struct {
	fx.In

	DbService database.Service
}

func New(p params) Service {
	return &restaurantService{
		db: p.DbService.GetDBInstance(),
	}
}

func (s restaurantService) CreateRestaurant(p CreateRestaurantParams) Restaurant {
	restaurant := Restaurant{
		Name:     p.Name,
		LogoPath: SaveImage(p.Logo),
	}
	s.db.Create(restaurant)
	return restaurant
}

func (s restaurantService) GetRestaurant(int64) Restaurant {
	var res models.Restaurant
	s.db.First(&res, 1)

	return Restaurant{
		Id:       res.Id,
		Name:     res.Name,
		LogoPath: res.LogoPath,
		Slogan:   res.Slogan,
	}
}

func (s restaurantService) UpdateQrCodeMenu(imgLink string, id string) string {
	outputFile := fmt.Sprintf("assets/qr-codes/%s_%s.png", id, imgLink)
	fmt.Println(outputFile)
	err := qrcode.WriteFile(imgLink, qrcode.Medium, 256, outputFile)
	fmt.Println(err)
	if err != nil {
		return "/"
	}
	return outputFile
}

func (s restaurantService) CreateBranch() {}
func (s restaurantService) GetBranches()  {}

// TODO(mazen): Add any saving image algorithm
func SaveImage(image *multipart.File) string {
	return "img"
}
