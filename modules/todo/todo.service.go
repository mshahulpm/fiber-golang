package todo

type TodoService struct{}

type Todo struct {
	ID     int    `json:"id"`
	Todo   string `json:"todo"`
	UserID int    `json:"userId"`
}

var todos = []Todo{}

func NewTodoService() *TodoService {
	return &TodoService{}
}

func (ts *TodoService) GetAllTodo() []Todo {
	return todos
}

func (ts *TodoService) AddNewTodo(todo Todo) Todo {

	todo.ID = len(todos) + 1
	todos = append(todos, todo)
	return todo
}
