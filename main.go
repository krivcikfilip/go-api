package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-api/database"
	"go-api/routes"
	"log"
)

// setupRoutes
// setup app routes
func setupRoutes(app *fiber.App) {
	api := app.Group("/api")

	routes.CreateCategoriesRoutes(api)
	routes.CreateBookRoutes(api)
}

// main func
func main() {
	database.ConnectDb()

	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
