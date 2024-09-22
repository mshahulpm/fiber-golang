package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	db "github.com/mshahulpm/todo-golang/database"
	"github.com/mshahulpm/todo-golang/modules/todo"
	"github.com/mshahulpm/todo-golang/modules/user"
)

func main() {

	db.InitDatabase()
	defer db.DB.Close()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		data := map[string]string{
			"message": "welcome to fiber server",
		}
		return c.JSON(data)
	})

	// load user module
	user.UserModule(app)

	// load todo module
	todo.TodoModule(app)

	fmt.Println("server is running on port 3000")
	app.Listen(":3000")
}
