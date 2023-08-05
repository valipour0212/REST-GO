package repository

import (
	"REST/database"
	"REST/model/user"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type UserRepository interface {
	GetUserList() ([]user.User, error)
	GetUserByID(id string) (user.User, error)
	GetUserByUserNameAndPassword(userName, password string) (user.User, error)
	InsertUser(user user.User) (string, error)
	UpdateUserByID(user user.User) error
	DeleteByID(id string) error
}

type userRepository struct {
	db database.DB
}

func NewUserRepository() UserRepository {
	db, err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	return userRepository{
		db: db,
	}
}

func (userRepository userRepository) GetUserList() ([]user.User, error) {

	userCollection := userRepository.db.GetUserCollection()

	cursor, err := userCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var users []user.User
	err = cursor.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}

	return users, nil

}

func (userRepository userRepository) GetUserByID(id string) (user.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user.User{}, nil
	}

	userCollection := userRepository.db.GetUserCollection()

	var userObject user.User
	err = userCollection.FindOne(context.TODO(), bson.D{
		{"_id", objectID},
	}).Decode(&userObject)

	if err != nil {
		return user.User{}, err
	}

	return userObject, nil
}

func (userRepository userRepository) GetUserByUserNameAndPassword(userName, password string) (user.User, error) {

	userCollection := userRepository.db.GetUserCollection()

	var userObject user.User
	err := userCollection.FindOne(context.TODO(), bson.D{
		{"UserName", userName},
		{"Password", password},
	}).Decode(&userObject)

	if err != nil {
		return user.User{}, err
	}

	return userObject, nil
}

func (userRepository userRepository) InsertUser(user user.User) (string, error) {
	userCollection := userRepository.db.GetUserCollection()

	result, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return "", err
	}

	objectID := result.InsertedID.(primitive.ObjectID).Hex()

	return objectID, nil
}

func (userRepository userRepository) UpdateUserByID(user user.User) error {
	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return err
	}

	user.ID = ""

	userCollection := userRepository.db.GetUserCollection()

	_, err = userCollection.UpdateOne(context.TODO(), bson.D{{"_id", objectID}}, bson.D{{"$set", user}})
	if err != nil {
		return err
	}

	return nil
}

func (userRepository userRepository) DeleteByID(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	userCollection := userRepository.db.GetUserCollection()

	_, err = userCollection.DeleteOne(context.TODO(), bson.D{{"_id", objectID}})
	if err != nil {
		return err
	}

	return nil
}
