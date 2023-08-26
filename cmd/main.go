package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pandakn/gofiber-crud/database"
	"github.com/pandakn/gofiber-crud/internal/routes"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	// Create a router group with "/api" prefix
	api := app.Group("/api")

	routes.SetupTodoRoutes(api)

	app.Listen(":8080")
}
