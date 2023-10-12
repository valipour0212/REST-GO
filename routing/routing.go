package routing

import (
	"REST/ViewModel/common/security"
	"REST/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetRouting(e *echo.Echo) error {

	e.POST("/login", controller.LoginUser)

	group := e.Group("users")

	group.GET("/getList", controller.GetUserList)
	jwtConfig := middleware.JWTConfig{
		SigningKey: []byte("secret"),
		Claims:     &security.JwtClaims{},
	}

	group.POST("/createNewUser", controller.CreateNewUser, middleware.JWTWithConfig(jwtConfig))

	return nil
}
