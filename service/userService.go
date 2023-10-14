package service

import (
	userViewModel "REST/ViewModel/user"
	"REST/model/user"
	"REST/repository"
	slices "slices"
	"time"
)

type UserService interface {
	CreateNewUser(userInput userViewModel.CreateNewUserViewModel) (string, error)
	//	GET
	GetUserList() ([]user.User, error)
	GetUserByUserNameAndPassword(loginViewModel userViewModel.LoginUserViewModel) (user.User, error)
	//	EDIT
	EditUser(userInput userViewModel.EditUserViewModel) error
	EditUserRole(userInput userViewModel.EditUserRoleViewModel) error
	EditUserPassword(userInput userViewModel.EditUserPasswordViewModel) error
	//	DELETE
	DeleteUser(id string) error
	//
	IsUserExist(id string) bool
	IsUserValidForAccess(userId, roleName string) bool
}

type userService struct {
}

func NewUserService() UserService {
	return userService{}
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

// GET
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

// EDIT
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
func (userService) EditUserRole(userInput userViewModel.EditUserRoleViewModel) error {
	userEntity := user.User{
		Id:    userInput.TargetUserId,
		Roles: userInput.Roles,
	}

	userRepository := repository.NewUserRepository()
	err := userRepository.UpdateUserById(userEntity)

	return err
}
func (userService) EditUserPassword(userInput userViewModel.EditUserPasswordViewModel) error {
	userEntity := user.User{
		Id:       userInput.TargetUserId,
		Password: userInput.Password,
	}

	userRepository := repository.NewUserRepository()
	err := userRepository.UpdateUserById(userEntity)

	return err
}

// DELETE
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
