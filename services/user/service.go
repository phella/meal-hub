package userService

import (
	"Bete/models"
	"Bete/services/database"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

type userService struct {
	db *gorm.DB
}

type params struct {
	fx.In

	DbService database.Service
}

func New(p params) Service {
	return &userService{
		db: p.DbService.GetDBInstance(),
	}
}

func (s userService) EnsureUser(p EnsureUserParams) models.User {
	user := models.User{
		Token: p.Token,
		Name:  p.Name,
	}
	s.db.Where(user).FirstOrCreate(&user)
	return user
}
