package routes

import (
	"github.com/gofiber/fiber/v2"
	todoHandlers "github.com/pandakn/gofiber-crud/internal/handlers"
)

func SetupTodoRoutes(api fiber.Router) {
	api.Post("/todos", todoHandlers.CreateTodo)
	api.Get("/todos", todoHandlers.GetTodos)
	api.Put("/todos/:id", todoHandlers.UpdateTodo)
	api.Delete("/todos/:id", todoHandlers.DeleteTodo)
}
