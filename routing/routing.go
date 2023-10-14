package routing

import (
	"REST/Utility"
	"REST/ViewModel/common/security"
	"REST/controller"
	"REST/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func SetRouting(e *echo.Echo) error {

	e.POST("/login", controller.LoginUser)

	group := e.Group("users")

	group.GET("/getList", controller.GetUserList)
	jwtConfig := middleware.JWTConfig{
		SigningKey: []byte("secret"),
		Claims:     &security.JwtClaims{},
	}

	group.POST("/createNewUser", controller.CreateNewUser, PermissionChecker("CreateUser"), middleware.JWTWithConfig(jwtConfig))

	group.PUT("/editUser/:id", controller.EditUser, PermissionChecker("EditUser"), middleware.JWTWithConfig(jwtConfig))

	group.DELETE("/deleteUser/:id", controller.DeleteUser, PermissionChecker("DeleteUser"), middleware.JWTWithConfig(jwtConfig))

	return nil
}

func PermissionChecker(permission string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			apiContext := c.(*Utility.ApiContext)

			operatorUserId, err := apiContext.GetUserId()
			if err != nil {
				return &echo.HTTPError{
					Code:     401,
					Message:  http.StatusUnauthorized,
					Internal: err,
				}
			}

			userService := service.NewUserService()
			isValid := userService.IsUserValidForAccess(operatorUserId, permission)
			if !isValid {
				return &echo.HTTPError{
					Code:    403,
					Message: http.StatusForbidden,
				}
			}

			return next(c)
		}
	}
}
