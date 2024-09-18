package user

import "github.com/gofiber/fiber/v2"

func UserModule(app *fiber.App) {
	// Initialize the service
	userService := NewUserService()

	// Initialize the controller
	userController := NewUserController(userService)

	// Setup routes
	userController.SetupRoutes(app)
}
