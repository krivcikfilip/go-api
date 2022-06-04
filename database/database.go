package database

import (
	"go-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var GORM *gorm.DB

// ConnectDb
// connect and initialize database with GORM
func ConnectDb() {

	// When using docker host={service_name} -> host=database
	// When using locally host=localhost
	dsn := "host=database user=root password=test dbname=api-db port=5432"

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
