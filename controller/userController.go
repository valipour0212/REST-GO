package controller

import (
	"REST/Utility"
	"REST/ViewModel/common/security"
	userViewModel "REST/ViewModel/user"
	"REST/service"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func GetUserList(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	fmt.Println(apiContext.GetUserId())

	userService := service.NewUserService()
	userList, err := userService.GetUserList()
	if err != nil {
		println(err)
	}

	return c.JSON(http.StatusOK, userList)
}

func CreateNewUser(c echo.Context) error {
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
		return c.JSON(http.StatusBadRequest, "")
	}

	if err := c.Validate(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, err)
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

	return c.JSON(http.StatusOK, userResData)
}

func EditUser(c echo.Context) error {
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

	return c.JSON(http.StatusOK, userResData)
}

func DeleteUser(c echo.Context) error {

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

	return c.JSON(http.StatusOK, userResData)
}

func LoginUser(c echo.Context) error {
	loginModel := new(userViewModel.LoginUserViewModel)

	if err := c.Bind(loginModel); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	if err := c.Validate(loginModel); err != nil {
		return c.JSON(http.StatusBadRequest, "Model not Valid")
	}

	userService := service.NewUserService()
	user, err := userService.GetUserByUserNameAndPassword(*loginModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "User Not found")
	}

	claims := &security.JwtClaims{
		UserName: user.UserName,
		UserId:   user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": stringToken,
	})
}
