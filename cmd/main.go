package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pandakn/go-fiber-simple-todo/database"
	"github.com/pandakn/go-fiber-simple-todo/internal/routes"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	// Create a router group with "/api" prefix
	api := app.Group("/api")

	routes.SetupTodoRoutes(api)

	app.Listen(":8080")
}
