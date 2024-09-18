package todo

import (
	"github.com/gofiber/fiber/v2"
)

type TodoController struct {
	service *TodoService
}

func NewTodoController(service *TodoService) *TodoController {
	return &TodoController{service: service}
}

func (uc *TodoController) SetupRoutes(app *fiber.App) {

	app.Get("/todo", func(c *fiber.Ctx) error {
		Todos := uc.service.GetAllTodo()
		return c.JSON(Todos)
	})

	app.Post("/todo", func(c *fiber.Ctx) error {

		var newTodo Todo

		// Parse request body into the Todo struct
		if err := c.BodyParser(&newTodo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		return c.JSON(uc.service.AddNewTodo(newTodo))
	})

}
