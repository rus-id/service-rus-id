package app

type UserService interface {
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}
