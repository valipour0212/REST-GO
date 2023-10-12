package controller

import (
	"REST/Utility"
	"REST/ViewModel/common/security"
	userViewModel "REST/ViewModel/user"
	"REST/service"
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

	newUser := new(userViewModel.CreateNewUserViewModel)

	if err := c.Bind(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	if err := c.Validate(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	creator, err := apiContext.GetUserId()
	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	newUser.CreatorUserId = creator

	userService := service.NewUserService()
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
