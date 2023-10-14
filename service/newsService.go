package service

import (
	"REST/model/news"
	"REST/repository"
)

type NewsService interface {
	GetNewsList() ([]news.News, error)
}

type newsService struct {
}

func NewNewsService() NewsService {
	return newsService{}
}

func (newsService) GetNewsList() ([]news.News, error) {

	newsRepository := repository.NewNewsRepository()
	newsList, err := newsRepository.GetNewsList()
	return newsList, err
}
