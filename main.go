package main

import (
	"REST/Utility"
	"REST/config"
	"REST/routing"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {

	//	Get Config
	err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Server port : ", config.AppConfig.Server.Port)

	//	Init Server
	server := echo.New()
	server.Validator = &Utility.CustomValidator{Validator: validator.New()}

	//	Routing
	routing.SetRouting(server)

	//	Middleware

	//	Start Server
	server.Start(":" + config.AppConfig.Server.Port)
}
