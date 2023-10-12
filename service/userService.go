package service

import (
	userViewModel "REST/ViewModel/user"
	"REST/model/user"
	"REST/repository"
)

type UserService interface {
	GetUserList() ([]user.User, error)
	CreateNewUser(userInput userViewModel.CreateNewUserViewModel) (string, error)
	GetUserByUserNameAndPassword(loginViewModel userViewModel.LoginUserViewModel) (user.User, error)
}

type userService struct {
}

func NewUserService() UserService {
	return userService{}
}

func (userService) GetUserList() ([]user.User, error) {

	userRepository := repository.NewUserRepository()
	userList, err := userRepository.GetUserList()

	return userList, err
}

func (userService) GetUserByUserNameAndPassword(loginViewModel userViewModel.LoginUserViewModel) (user.User, error) {

	userRepository := repository.NewUserRepository()
	user, err := userRepository.GetUserByUserNameAndPassword(loginViewModel.UserName, loginViewModel.Password)

	return user, err
}

func (userService) CreateNewUser(userInput userViewModel.CreateNewUserViewModel) (string, error) {

	userEntity := user.User{}

	userRepository := repository.NewUserRepository()
	userId, err := userRepository.InsertUser(userEntity)

	return userId, err
}
