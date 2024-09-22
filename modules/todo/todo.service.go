package todo

import (
	"context"
	"fmt"

	db "github.com/mshahulpm/todo-golang/database"
	"github.com/uptrace/bun"
)

type TodoService struct{}

type ITodo struct {
	TodoId int64
	Todo   string
	UserID string
}
type Todo struct {
	ID     int    `json:"id"`
	Todo   string `json:"todo"`
	UserID int    `json:"userId"`
}

var todos = []Todo{}

func NewTodoService() *TodoService {
	return &TodoService{}
}

func (ts *TodoService) GetAllTodo() []ITodo {

	todooos := make([]ITodo, 0)

	err := db.DB.NewRaw(
		"SELECT todo_id,user_id,todo from ? limit ?",
		bun.Ident("todo"),
		100,
	).Scan(context.Background(), &todooos)

	if err != nil {
		fmt.Println(err.Error())
	}

	return todooos
}

func (ts *TodoService) AddNewTodo(todo Todo) Todo {

	todo.ID = len(todos) + 1
	todos = append(todos, todo)
	return todo
}
