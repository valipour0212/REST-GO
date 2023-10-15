package service

import (
	newsViewModel "REST/ViewModel/news"
	"REST/model/news"
	"REST/repository"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type NewsService interface {
	GetNewsList() ([]news.News, error)
	CreateNewUser(userInput newsViewModel.CreateNewsViewModel, imageFile *multipart.FileHeader) (string, error)
}

type newsService struct {
}

func NewNewsService() NewsService {
	return newsService{}
}

// ---------------------------------------------------
func (newsService) GetNewsList() ([]news.News, error) {

	newsRepository := repository.NewNewsRepository()
	newsList, err := newsRepository.GetNewsList()
	return newsList, err
}

func (NewsSer newsService) CreateNewUser(userInput newsViewModel.CreateNewsViewModel, imageFile *multipart.FileHeader) (string, error) {

	newsEntity := news.News{
		Title:            userInput.Title,
		ImageName:        userInput.ImageName,
		ShortDescription: userInput.ShortDescription,
		Description:      userInput.Description,
		CreateDate:       time.Now(),
		CreatorUserId:    userInput.CreatorUserId,
	}
	if imageFile != nil {
		src, err := imageFile.Open()
		if err != nil {
			return "", err
		}

		fileName := uuid.New().String() + filepath.Ext(imageFile.Filename)

		wd, err := os.Getwd()
		imageServerPath := filepath.Join(wd, "wwwRoot", "images", "news", fileName)

		des, err := os.Create(imageServerPath)
		if err != nil {
			return "", err
		}
		defer des.Close()

		_, err = io.Copy(des, src)
		if err != nil {
			return "", err
		}
		newsEntity.ImageName = fileName
	}

	newsRepository := repository.NewNewsRepository()
	newsId, err := newsRepository.InsertNews(newsEntity)

	return newsId, err
}
