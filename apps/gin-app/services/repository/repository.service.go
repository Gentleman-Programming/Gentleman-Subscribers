package service

import (
	models "apps/gin-app/models"
	db_config "apps/gin-app/services/repository/config"
	subscriberModel "apps/gin-app/services/subscribers/models"
	"context"
	"log/slog"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
)

func InitializeDb() error {
	_, err := db_config.InitializeDb()
	if err != nil {
		slog.Error(err.Message)

		return err.Err
	}

	return nil
}

type RepositoryService struct{}

func NewRepositoryService() (*RepositoryService, error) {
	err := InitializeDb()
	if err != nil {
		return nil, err
	}

	return &RepositoryService{}, nil
}

func (n *RepositoryService) saveToDatabase(subscriber subscriberModel.Subscriber) *models.AppError {
	_, err := db_config.SubscriberCollection.InsertOne(context.TODO(), subscriber)
	if err != nil {
		return &models.AppError{
			Err:     err,
			Message: "failed to save subscriber",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

func (n *RepositoryService) SetSubscribers(subscribers *[]subscriberModel.Subscriber) *models.AppError {
	newSubscribers := make([]interface{}, len(*subscribers))
	for i, subscriber := range *subscribers {
		go n.saveToDatabase(subscriber)
		newSubscribers[i] = subscriber
	}

	_, err := db_config.SubscriberCollection.InsertMany(context.TODO(), newSubscribers)
	if err != nil {
		return &models.AppError{
			Err:     err,
			Message: "failed to set subscribers",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}

func (n *RepositoryService) GetAllSubscribers() (*[]subscriberModel.Subscriber, *models.AppError) {
	var result []subscriberModel.Subscriber

	cursor, err := db_config.SubscriberCollection.Find(context.TODO(), bson.D{}, options.Find())
	if err != nil {
		return nil, &models.AppError{
			Err:     err,
			Message: "failed to get all subscribers",
			Code:    http.StatusInternalServerError,
		}
	}

	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &result); err != nil {
		return nil, &models.AppError{
			Err:     err,
			Message: "failed to decode subscribers",
			Code:    http.StatusInternalServerError,
		}
	}

	for _, element := range result {
		_, err := bson.MarshalExtJSON(element, false, false)
		if err != nil {
			return nil, &models.AppError{
				Err:     err,
				Message: "Failed Unmarshal BSON",
				Code:    http.StatusInternalServerError,
			}
		}
	}

	if err := cursor.Err(); err != nil {
		return nil, &models.AppError{
			Err:     err,
			Message: "Mongo Cursor Error",
			Code:    http.StatusInternalServerError,
		}
	}
	return &result, nil
}

var Module = fx.Options(
	fx.Provide(NewRepositoryService),
)
