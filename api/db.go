package api

import (
	"fmt"
	"log"
	"os"

	"github.com/Shreeyash-Naik/Valet-Parking/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDB() {
	if Db != nil {
		return
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to DB")
	}

	Db = db
	log.Println("Successfully connected to DB")
}

func Migrate() {
	Db.AutoMigrate(&models.Vehicle{})
}
