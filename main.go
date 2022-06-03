package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-api/database"
	"go-api/routes"
	"log"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")

	routes.CreateCategoriesRoutes(api)
	routes.CreateBookRoutes(api)
}

func main() {
	database.ConnectDb()

	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen("localhost:3000"))
}
