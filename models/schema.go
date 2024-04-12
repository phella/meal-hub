package models

import (
	"Bete/services/database"
	"go.uber.org/fx"
)

type migrateSchemaParams struct {
	fx.In

	DbService database.Service
}

func MigrateSchema(params migrateSchemaParams) error {
	db := params.DbService.GetDBInstance()
	if err := db.AutoMigrate(&Restaurant{}, &Dish{}); err != nil {
		return err
	}

	return nil
}
