package main

import (
	"github.com/gofiber/fiber/v2"
	"go-api/database"
	"go-api/routes"
	"log"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/categories", routes.FindCategories)
}

func main() {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
