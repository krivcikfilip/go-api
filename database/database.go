package database

import (
	"go-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var GORM *gorm.DB

func ConnectDb() {
	dsn := "host=localhost user=root password=test dbname=api-db port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error with connection to database", err)
	}

	err = db.AutoMigrate(&models.Category{}, &models.Book{})
	if err != nil {
		log.Println("Error with migrations", err)
	}

	GORM = db
}
