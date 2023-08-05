package controller

import (
	"REST/service"
	"REST/viewModel/common/security"
	userViewModel "REST/viewModel/user"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func GetUserList(ctx echo.Context) error {

	userService := service.NewUserService()
	userList, err := userService.GetUserList()
	if err != nil {
		println(err)
	}

	return ctx.JSON(http.StatusOK, userList)
}

func CreateNewUser(ctx echo.Context) error {
	newUser := new(userViewModel.CreateNewUserViewModel)

	if err := ctx.Bind(newUser); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := ctx.Validate(newUser); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	token := ctx.Get("user").(*jwt.Token)
	claims := token.Claims.(*security.JWTClaims)
	newUser.CreatorUserId = claims.UserID

	userService := service.NewUserService()
	newUserID, err := userService.CreateNewUser(*newUser)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	userResponseData := struct {
		NewUserID string
	}{
		NewUserID: newUserID,
	}
	return ctx.JSON(http.StatusOK, userResponseData)
}

func LoginUser(ctx echo.Context) error {
	loginModel := new(userViewModel.LoginUserViewModel)

	if err := ctx.Bind(loginModel); err != nil {
		return ctx.JSON(http.StatusBadRequest, "")
	}

	if err := ctx.Validate(loginModel); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Model Not Valid")
	}

	userService := service.NewUserService()
	user, err := userService.GetUserByUserNameAndPassword(*loginModel)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Not Found User")
	}

	claims := &security.JWTClaims{
		UserName: user.UserName,
		UserID:   user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	stringToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	userResData := struct {
		Token string
	}{
		Token: stringToken,
	}

	return ctx.JSON(http.StatusOK, userResData)
}
