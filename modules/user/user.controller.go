package user

import (
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	service *UserService
}

func NewUserController(service *UserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) GetUsers(c *fiber.Ctx) error {
	users := uc.service.GetAllUsers()
	return c.JSON(users)
}

func (uc *UserController) SetupRoutes(app *fiber.App) {
	app.Get("/users", uc.GetUsers)
}
