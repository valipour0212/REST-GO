package main

import (
	"REST/Utility"
	"REST/ViewModel/common/security"
	"REST/config"
	"REST/routing"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func RootLevel(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("RootLevel")
	return next
}
func AfterRouter(next echo.HandlerFunc) echo.HandlerFunc {
	fmt.Println("AfterRouter")
	return next
}

func main() {

	// config
	err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Server port : ", config.AppConfig.Server.Port)

	//init server
	server := echo.New()
	server.Validator = &Utility.CustomValidator{Validator: validator.New()}

	//routing
	routing.SetRouting(server)

	//middleware
	server.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			apiContext := &Utility.ApiContext{Context: c}

			return next(apiContext)
		}
	})
	server.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:             []byte("secret"),
		Claims:                 &security.JwtClaims{},
		ContinueOnIgnoredError: true,
		ErrorHandlerWithContext: func(err error, c echo.Context) error {
			return nil
		},
	}))
	server.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(1)))

	//start server
	server.Start(":" + config.AppConfig.Server.Port)
}
