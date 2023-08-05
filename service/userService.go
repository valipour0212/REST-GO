package service

import (
	"REST/model/user"
	"REST/repository"
	userViewModel "REST/viewModel/user"
	"time"
)

type UserService interface {
	GetUserList() ([]user.User, error)
	CreateNewUser(userModel userViewModel.CreateNewUserViewModel) (string, error)
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

func (userService) CreateNewUser(userModel userViewModel.CreateNewUserViewModel) (string, error) {
	userEntity := user.User{
		FirstName:     userModel.FirstName,
		LastName:      userModel.LastName,
		Email:         userModel.Email,
		UserName:      userModel.UserName,
		Password:      userModel.Password,
		RegisterDate:  time.Now(),
		CreatorUserID: userModel.CreatorUserId,
	}

	userRepository := repository.NewUserRepository()

	userID, err := userRepository.InsertUser(userEntity)
	if err != nil {
		return "", err
	}

	return userID, nil
}
