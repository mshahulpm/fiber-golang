package user

import (
	"context"
	"fmt"
	"log"

	db "github.com/mshahulpm/todo-golang/database"
	models "github.com/mshahulpm/todo-golang/model"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) GetAllUsers() []models.UserModel {

	users := make([]models.UserModel, 1)

	err := db.DB.NewRaw(
		"select user_id,email,created_at,updated_at from ? limit ?",
		bun.Ident("users"),
		100,
	).Scan(context.Background(), &users)

	if err != nil {
		fmt.Println(err.Error())
	}

	return users

}

func (us *UserService) SignUp(email, password string) error {
	// Hash the password before saving it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create a map with the user values
	values := map[string]interface{}{
		"email":    email,
		"password": string(hashedPassword), // Save the hashed password
	}

	// Insert the new user into the users table
	_, err = db.DB.NewInsert().
		Model(&values).
		TableExpr("users"). // Specify the table name
		Exec(context.Background())
	if err != nil {
		log.Println("Error inserting new user:", err)
		return err
	}

	return nil
}
