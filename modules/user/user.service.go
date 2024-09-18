package user

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) GetAllUsers() []string {
	return []string{"John", "Jane", "Alice"}
}
