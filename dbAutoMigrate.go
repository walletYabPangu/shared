package shared

import (
	"github.com/walletYabPangu/shared/models"
	"gorm.io/gorm"
)

var db *gorm.DB
var initializers = []func() error{
	routeCreate,
}

func routeCreate() error {
	err := db.AutoMigrate(models.ServiceRoute{})

	if err != nil {
		var mR = []models.ServiceRoute{
			{
				ServiceKey:  "auth",
				UpstreamURL: "http://localhost:8081",
			},
			{
				ServiceKey:  "user",
				UpstreamURL: "http://localhost:8082",
			},
			{
				ServiceKey:  "game",
				UpstreamURL: "http://localhost:8083",
			},
			{
				ServiceKey:  "task",
				UpstreamURL: "http://localhost:8084",
			},
			{
				ServiceKey:  "shop",
				UpstreamURL: "http://localhost:8085",
			},
			{
				ServiceKey:  "admin",
				UpstreamURL: "http://localhost:8086",
			},
		}

		db.Save(mR)
	}

	return err
}

func (d *Database) ConnectAndMigrate() error {
	db = d.Db
	for _, initialize := range initializers {
		if err := initialize(); err != nil {
			return err
		}
	}
	return nil
}
