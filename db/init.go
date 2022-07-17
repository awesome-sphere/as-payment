package db

import (
	"fmt"
	"log"

	"github.com/awesome-sphere/as-payment/models"
	"github.com/awesome-sphere/as-payment/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() {
	dbUser := utils.GetenvOr("POSTGRES_USER", "pkinwza")
	dbPassword := utils.GetenvOr("POSTGRES_PASSWORD", "securepassword")
	dbHost := utils.GetenvOr("POSTGRES_HOST", "localhost")
	dbPort := utils.GetenvOr("POSTGRES_PORT", "5432")
	dbName := utils.GetenvOr("POSTGRES_DB", "as-cinema") // FIXME: change the database when we make one

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Order{}, &models.OrderSeats{})

	DB = db
}
