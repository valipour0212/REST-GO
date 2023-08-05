package routing

import (
	"REST/controller"
	"REST/viewModel/common/security"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetRouting(e *echo.Echo) error {

	e.POST("/login", controller.LoginUser)

	Group := e.Group("users")

	Group.GET("/getList", controller.GetUserList)

	jwtConfig := middleware.JWTConfig{
		SigningKey: []byte("secret"),
		Claims:     &security.JWTClaims{},
	}

	Group.POST("/createNewUser", controller.CreateNewUser, middleware.JWTWithConfig(jwtConfig))

	return nil
}
