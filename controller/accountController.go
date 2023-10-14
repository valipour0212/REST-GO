package controller

import (
	"REST/ViewModel/common/security"
	userViewModel "REST/ViewModel/user"
	"REST/service"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type AccountController interface {
	LoginUser(c echo.Context) error
}

type accountController struct {
	userService service.UserService
}

func NewAccountController() AccountController {
	userService := service.NewUserService()
	return accountController{
		userService: userService,
	}
}

func (AccountController accountController) LoginUser(c echo.Context) error {
	loginModel := new(userViewModel.LoginUserViewModel)

	if err := c.Bind(loginModel); err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}

	if err := c.Validate(loginModel); err != nil {
		return c.JSON(http.StatusBadRequest, "Model not Valid")
	}

	user, err := AccountController.userService.GetUserByUserNameAndPassword(*loginModel)
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
