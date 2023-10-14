package controller

import (
	"REST/Utility"
	"REST/ViewModel/common/httpResponse"
	"REST/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type NewsController interface {
	GetNewsList(c echo.Context) error
}

type newsController struct {
}

func NewNewsController() NewsController {
	return newsController{}
}

func (NewsC newsController) GetNewsList(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	fmt.Println(apiContext.GetUserId())

	newsService := service.NewNewsService()
	newsList, err := newsService.GetNewsList()
	if err != nil {
		println(err)
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(newsList))
}
