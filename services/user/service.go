package userService

import (
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

func (s userService) CreateUser(p CreateUserParams) User {
	var user User 
	s.db.Where(User{Token: p.Token}).FirstOrCreate(&user)
	return user
}
