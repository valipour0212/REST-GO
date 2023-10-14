package controller

import (
	"REST/Utility"
	"REST/ViewModel/common/httpResponse"
	userViewModel "REST/ViewModel/user"
	"REST/service"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type UserController interface {
	CreateNewUser(c echo.Context) error

	//	GET
	GetUserList(c echo.Context) error

	//	EDIT
	EditUser(c echo.Context) error
	EditUserRole(c echo.Context) error
	EditUserPassword(c echo.Context) error

	//	DELETE
	DeleteUser(c echo.Context) error

	//
	UploadAvatar(c echo.Context) error
}

type userController struct {
}

func NewUserController() UserController {
	return userController{}
}

func (UC userController) CreateNewUser(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)

	operatorUserId, err := apiContext.GetUserId()
	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	userService := service.NewUserService()

	isValid := userService.IsUserValidForAccess(operatorUserId, "")
	if !isValid {
		return c.JSON(http.StatusForbidden, "")
	}

	newUser := new(userViewModel.CreateNewUserViewModel)

	if err := c.Bind(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse("Data Not Found"))
	}

	if err := c.Validate(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse(err))
	}

	newUser.CreatorUserId = operatorUserId

	newUserId, err := userService.CreateNewUser(*newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		NewUserId string
	}{
		NewUserId: newUserId,
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(userResData))
}

// GET
func (UC userController) GetUserList(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	fmt.Println(apiContext.GetUserId())

	userService := service.NewUserService()
	userList, err := userService.GetUserList()
	if err != nil {
		println(err)
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(userList))
}

// EDIT
func (UC userController) EditUser(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)

	targetUserID := apiContext.Param("id")
	fmt.Println(targetUserID)

	userService := service.NewUserService()

	newUserData := new(userViewModel.EditUserViewModel)

	if err := c.Bind(newUserData); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	if err := c.Validate(newUserData); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	newUserData.TargetUserID = targetUserID

	if !userService.IsUserExist(targetUserID) {
		return c.JSON(http.StatusBadRequest, errors.New("Not Found User"))
	}

	err := userService.EditUser(*newUserData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		IsSuccess bool
	}{
		IsSuccess: true,
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(userResData))
}
func (UC userController) EditUserRole(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	targetUserId := apiContext.Param("id")

	userService := service.NewUserService()
	newUserData := new(userViewModel.EditUserRoleViewModel)

	if err := c.Bind(newUserData); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	if err := c.Validate(newUserData); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	newUserData.TargetUserId = targetUserId

	if !userService.IsUserExist(targetUserId) {
		return c.JSON(http.StatusBadRequest, errors.New("User Not Found"))
	}

	err := userService.EditUserRole(*newUserData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		IsSuccess bool
	}{
		IsSuccess: true,
	}

	return c.JSON(http.StatusOK, userResData)
}
func (UC userController) EditUserPassword(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	targetUserId := apiContext.Param("id")

	userService := service.NewUserService()
	newUserData := new(userViewModel.EditUserPasswordViewModel)

	if err := c.Bind(newUserData); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	if err := c.Validate(newUserData); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	newUserData.TargetUserId = targetUserId

	if !userService.IsUserExist(targetUserId) {
		return c.JSON(http.StatusBadRequest, errors.New("User Not Found"))
	}

	err := userService.EditUserPassword(*newUserData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		IsSuccess bool
	}{
		IsSuccess: true,
	}

	return c.JSON(http.StatusOK, userResData)
}

// DELETE
func (UC userController) DeleteUser(c echo.Context) error {

	apiContext := c.(*Utility.ApiContext)

	targetUserId := apiContext.Param("id")

	userService := service.NewUserService()
	if !userService.IsUserExist(targetUserId) {
		return c.JSON(http.StatusBadRequest, errors.New("User Not Found"))
	}

	err := userService.DeleteUser(targetUserId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		IsSuccess bool
	}{
		IsSuccess: true,
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(userResData))
}

func (UC userController) UploadAvatar(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)

	file, err := apiContext.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	wd, err := os.Getwd()
	imageServerPath := filepath.Join(wd, "wwwRoot", "images", "userAvatar", file.Filename)

	des, err := os.Create(imageServerPath)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	defer des.Close()

	_, err = io.Copy(des, src)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		IsSuccess bool
	}{
		IsSuccess: true,
	}

	return c.JSON(http.StatusOK, userResData)
}
