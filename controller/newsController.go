package controller

import (
	"REST/Utility"
	"REST/ViewModel/common/httpResponse"
	newsViewModel "REST/ViewModel/news"
	"REST/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type NewsController interface {
	GetNewsList(c echo.Context) error
	CreateNews(c echo.Context) error
}

type newsController struct {
}

func NewNewsController() NewsController {
	return newsController{}
}

// -----------------------------------------------------
func (NewsCon newsController) GetNewsList(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	fmt.Println(apiContext.GetUserId())

	newsService := service.NewNewsService()
	newsList, err := newsService.GetNewsList()
	if err != nil {
		println(err)
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(newsList))
}

func (NewsCon newsController) CreateNews(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)

	newNews := new(newsViewModel.CreateNewsViewModel)

	if err := apiContext.Bind(newNews); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse("Data not found"))
	}

	if err := c.Validate(newNews); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse(err))
	}

	file, err := apiContext.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse("image not found"))
	}

	newsService := service.NewNewsService()
	newNewsId, err := newsService.CreateNewUser(*newNews, file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		NewUserId string
	}{
		NewUserId: newNewsId,
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(userResData))
}
