package todo

import "github.com/gofiber/fiber/v2"

func TodoModule(app *fiber.App) {
	// Initialize the service
	todoService := NewTodoService()

	// Initialize the controller
	todoController := NewTodoController(todoService)

	// Setup routes
	todoController.SetupRoutes(app)
}
