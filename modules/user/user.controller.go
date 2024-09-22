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

func (uc *UserController) SignUpHandler(c *fiber.Ctx) error {

	// Extract data from request body (assuming JSON)
	type SignUpRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req SignUpRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Call the SignUp method
	err := uc.service.SignUp(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}

func (uc *UserController) SetupRoutes(app *fiber.App) {
	app.Get("/users", uc.GetUsers)
	app.Post("/signup", uc.SignUpHandler)
}
