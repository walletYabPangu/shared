package shared

import (
	"fmt"
	"github.com/walletYabPangu/shared/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Database struct {
	Db *gorm.DB
}

type IDatabase interface {
	GetRoutes() (*models.ServiceRoute, error)
	ConnectAndMigrate() error
}

func InitDb(Data *DbConfig) IDatabase {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", Data.Host, Data.User, Data.Password, Data.DBName, Data.Port)
	dbi, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	return &Database{
		Db: dbi,
	}
}

func (d *Database) ConnectAndMigrate() error {
	er := d.Db.AutoMigrate(models.ServiceRoute{})
	return er
}

func (d *Database) GetRoutes() (*models.ServiceRoute, error) {

	r := models.ServiceRoute{}

	err := d.Db.Model(models.ServiceRoute{}).Find(&r).Error
	if err != nil {
		return nil, err
	}

	return &r, err
}

//func (d *Database) SetRoutes(route *models.ServiceRoute) error {
//	d.Db.Model(models.ServiceRoute{}).Where("")
//}
