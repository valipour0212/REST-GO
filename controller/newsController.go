package controller

import (
	"REST/Utility"
	"REST/ViewModel/common/httpResponse"
	newsViewModel "REST/ViewModel/news"
	"REST/service"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type NewsController interface {
	GetNewsList(c echo.Context) error
	GetNews(c echo.Context) error
	CreateNews(c echo.Context) error
	EditNews(c echo.Context) error
	DeleteNews(c echo.Context) error
	LikeNews(c echo.Context) error
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

func (NewsCon newsController) GetNews(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	targetNewsId := apiContext.Param("id")

	newsService := service.NewNewsService()
	news, err := newsService.GetNewsById(targetNewsId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.NotFoundResponse(nil, "news not found"))
	}

	newsService.AddVisitCount(targetNewsId)

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(news))
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
	newNewsId, err := newsService.CreateNews(*newNews, file)
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

func (NewsCon newsController) EditNews(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	targetNewsId := apiContext.Param("id")

	editNews := new(newsViewModel.EditNewsViewModel)

	if err := apiContext.Bind(editNews); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse("Data not found"))
	}

	if err := c.Validate(editNews); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse(err))
	}

	file, err := apiContext.FormFile("file")

	editNews.Id = targetNewsId

	newsService := service.NewNewsService()

	if !newsService.IsNewsExist(targetNewsId) {
		return c.JSON(http.StatusBadRequest, errors.New("User Not Found"))
	}

	err = newsService.EditNews(*editNews, file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(nil))
}

func (NewsCon newsController) DeleteNews(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	targetNewsId := apiContext.Param("id")

	newsService := service.NewNewsService()
	if !newsService.IsNewsExist(targetNewsId) {
		return c.JSON(http.StatusNotFound, httpResponse.SuccessResponse("Data not found"))
	}

	err := newsService.DeleteNews(targetNewsId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(nil))
}

func (NewsCon newsController) LikeNews(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	targetNewsId := apiContext.Param("id")

	newsService := service.NewNewsService()

	if !newsService.IsNewsExist(targetNewsId) {
		return c.JSON(http.StatusBadRequest, errors.New("User Not Found"))
	}

	err := newsService.AddLike(targetNewsId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(nil))
}
