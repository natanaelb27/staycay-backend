package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/natanaelb27/staycay-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() (err error) {
	// dsn := fmt.Sprintf(
	// 	"host=%v port=%v user=%v password=%v dbname=%v sslmode=require",
	// 	os.Getenv("HOST"),
	// 	os.Getenv("PORT"),
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASS"),
	// 	os.Getenv("DB_NAME"))

	dsn := ""
	if os.Getenv("ENV") == "deployment" {
		dsn = os.Getenv("DATABASE_URL")
	} else {
		godotenv.Load(".env")
		dsn = fmt.Sprintf(
			"host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
			os.Getenv("HOST"),
			os.Getenv("PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME"))
	}

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.City{}, &models.User{}, &models.House{}, &models.Wallet{}, &models.GameChance{}, &models.HousePhoto{}, &models.Reservation{}, &models.PickupStatus{}, &models.Pickup{})

	if err != nil {
		panic(err)
	}

	return nil
}

func Get() *gorm.DB {
	return db
}
