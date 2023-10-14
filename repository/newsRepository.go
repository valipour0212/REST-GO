package repository

import (
	"REST/database"
	"REST/model/news"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type NewsRepository interface {
	InsertNews(user news.News) (string, error)

	//	GET
	GetNewsList() ([]news.News, error)
	GetNewsById(id string) (news.News, error)

	//	Edit
	UpdateNewsById(user news.News) error

	//	DELETE
	DeleteNewsById(id string) error
}

type newsRepository struct {
	db database.Db
}

func NewNewsRepository() NewsRepository {
	db, err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	return newsRepository{
		db: db,
	}
}

func (NewsRep newsRepository) InsertNews(news news.News) (string, error) {

	newsCollection := NewsRep.db.GetNewsCollection()
	res, err := newsCollection.InsertOne(context.TODO(), news)

	if err != nil {
		return "", err
	}
	objectId := res.InsertedID.(primitive.ObjectID).Hex()
	return objectId, nil
}

// GET
func (NewsRep newsRepository) GetNewsList() ([]news.News, error) {

	newsCollection := NewsRep.db.GetNewsCollection()

	cursor, err := newsCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var news []news.News
	err = cursor.All(context.TODO(), &news)
	if err != nil {
		return nil, err
	}

	return news, nil

}
func (NewsRep newsRepository) GetNewsById(id string) (news.News, error) {

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return news.News{}, err
	}
	newsCollection := NewsRep.db.GetNewsCollection()
	var userObject news.News
	// db.getCollection('users').find({"_id" : ObjectId("6297bdcc8d7757574658ed66")})
	err = newsCollection.FindOne(context.TODO(), bson.D{
		{"_id", objectId},
	}).Decode(&userObject)

	if err != nil {
		return news.News{}, err
	}

	return userObject, nil

}

// Edit
func (NewsRep newsRepository) UpdateNewsById(news news.News) error {
	objectId, err := primitive.ObjectIDFromHex(news.Id)
	if err != nil {
		return err
	}
	news.Id = ""
	newsCollection := NewsRep.db.GetNewsCollection()
	_, err = newsCollection.UpdateOne(context.TODO(), bson.D{{"_id", objectId}}, bson.D{{"$set", news}})

	if err != nil {
		return err
	}

	return nil
}

// DELETE
func (NewsRep newsRepository) DeleteNewsById(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	newsCollection := NewsRep.db.GetNewsCollection()
	_, err = newsCollection.DeleteOne(context.TODO(), bson.D{{"_id", objectId}})

	if err != nil {
		return err
	}

	return nil
}
