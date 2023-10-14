package service

import (
	userViewModel "REST/ViewModel/user"
	"REST/model/user"
	"REST/repository"
	slices "slices"
	"time"
)

type UserService interface {
	GetUserList() ([]user.User, error)
	CreateNewUser(userInput userViewModel.CreateNewUserViewModel) (string, error)
	GetUserByUserNameAndPassword(loginViewModel userViewModel.LoginUserViewModel) (user.User, error)
	EditUser(userInput userViewModel.EditUserViewModel) error
	DeleteUser(id string) error
	IsUserExist(id string) bool
	IsUserValidForAccess(userId, roleName string) bool
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

func (userService) IsUserValidForAccess(userId, roleName string) bool {

	userRepository := repository.NewUserRepository()
	user, err := userRepository.GetUserById(userId)
	if err != nil {
		return false
	}

	if user.Roles != nil {
		return false
	}

	roleIndex := slices.IndexFunc(user.Roles, func(role string) bool {
		return role == roleName
	})
	return roleIndex >= 0
}

func (userService) CreateNewUser(userInput userViewModel.CreateNewUserViewModel) (string, error) {

	userEntity := user.User{
		FirstName:     userInput.FirstName,
		LastName:      userInput.LastName,
		Email:         userInput.Email,
		UserName:      userInput.UserName,
		Password:      userInput.Password,
		RegisterDate:  time.Time{},
		CreatorUserId: userInput.CreatorUserId,
	}

	userRepository := repository.NewUserRepository()
	userId, err := userRepository.InsertUser(userEntity)

	return userId, err
}

func (userService) EditUser(userInput userViewModel.EditUserViewModel) error {

	userEntity := user.User{
		Id:           userInput.TargetUserID,
		FirstName:    userInput.FirstName,
		LastName:     userInput.LastName,
		Email:        userInput.Email,
		UserName:     userInput.UserName,
		Password:     userInput.Password,
		RegisterDate: time.Time{},
	}

	userRepository := repository.NewUserRepository()
	err := userRepository.UpdateUserById(userEntity)

	return err
}

func (userService) DeleteUser(id string) error {

	userRepository := repository.NewUserRepository()
	err := userRepository.DeleteUserById(id)

	return err
}

func (userService) IsUserExist(id string) bool {

	userRepository := repository.NewUserRepository()
	_, err := userRepository.GetUserById(id)
	if err != nil {
		return false
	}

	return true
}
