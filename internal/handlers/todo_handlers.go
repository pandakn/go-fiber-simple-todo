package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/pandakn/go-fiber-simple-todo/database"
	"github.com/pandakn/go-fiber-simple-todo/internal/models"
)

func GetTodos(c *fiber.Ctx) error {
	todos := []models.Todo{}
	database.DB.Db.Order("id").Find(&todos)

	return c.JSON(fiber.Map{"result": todos})
}

func CreateTodo(c *fiber.Ctx) error {
	var newTodo models.Todo

	fmt.Print("newTodo:", newTodo)

	if err := c.BodyParser(&newTodo); err != nil {
		return err
	}

	// Set empty title to null
	if newTodo.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Title is required"})
	}

	database.DB.Db.Create(&newTodo)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "created successfully", "result": newTodo})
}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var updatedTodo models.Todo
	if err := c.BodyParser(&updatedTodo); err != nil {
		return err
	}

	var todo models.Todo
	if err := database.DB.Db.First(&todo, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	}

	database.DB.Db.Model(&todo).Updates(updatedTodo)

	return c.JSON(fiber.Map{"message": "created successfully", "result": todo})
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	var todo models.Todo

	if err := database.DB.Db.First(&todo, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Todo not found"})
	}

	database.DB.Db.Delete(&todo)

	return c.JSON(fiber.Map{"message": "deleted successfully"})
}
