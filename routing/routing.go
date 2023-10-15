package routing

import (
	customMiddlewares "REST/Utility/middleware"
	"REST/config"
	"REST/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetRouting(e *echo.Echo) error {

	RouteUserController(e)
	RouteAccountController(e)
	RouteNewsController(e)

	return nil
}

func RouteUserController(e *echo.Echo) {
	userController := controller.NewUserController()

	e.POST("/UploadAvatar", userController.UploadAvatar)

	userGroup := e.Group("users")

	userGroup.GET("/getList", userController.GetUserList)

	userGroup.POST("/CreateNews", userController.CreateNewUser, customMiddlewares.PermissionChecker("CreateUser"), middleware.JWTWithConfig(config.AppConfig.DefJwtConfig))

	userGroup.PUT("/EditUser/:id", userController.EditUser, customMiddlewares.PermissionChecker("EditUser"), middleware.JWTWithConfig(config.AppConfig.DefJwtConfig))
	userGroup.DELETE("/DeleteUser/:id", userController.DeleteUser, customMiddlewares.PermissionChecker("DeleteUser"), middleware.JWTWithConfig(config.AppConfig.DefJwtConfig))

	userGroup.PUT("/EditUserRole/:id", userController.EditUserRole, middleware.JWTWithConfig(config.AppConfig.DefJwtConfig))
	userGroup.PUT("/EditUserPassword/:id", userController.EditUserPassword, middleware.JWTWithConfig(config.AppConfig.DefJwtConfig))

}

func RouteAccountController(e *echo.Echo) {
	accountController := controller.NewAccountController()
	e.POST("/login", accountController.LoginUser)
}

func RouteNewsController(e *echo.Echo) {
	newsController := controller.NewNewsController()

	newsGroup := e.Group("news")
	newsGroup.GET("/getList", newsController.GetNewsList)
	newsGroup.POST("/create", newsController.CreateNews)
	newsGroup.POST("/edit/:id", newsController.EditNews)
}
