package db

import (
	"database/sql"

	models "github.com/mshahulpm/todo-golang/model"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var DB *bun.DB
var User *models.UserModel
var Todo *models.TodoModel

func InitDatabase() {

	dsn := "postgres://postgres:123456@localhost:5432/golang_todo?sslmode=disable"
	// dsn := "unix://user:pass@dbname/var/run/postgresql/.s.PGSQL.5432"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	DB = bun.NewDB(sqldb, pgdialect.New())

	User = new(models.UserModel)
	Todo = new(models.TodoModel)
}
