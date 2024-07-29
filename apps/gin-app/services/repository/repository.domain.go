package service

import (
	"apps/gin-app/models"
	db_config "apps/gin-app/services/repository/config"
	subscriberModel "apps/gin-app/services/subscribers/models"
	"context"
	"log/slog"
	"net/http"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitializeDb() {
	_, err := db_config.InitializeDb()
	if err != nil {
		slog.Error(err.Message)

		panic(err.Err)
	}
}

type SubscriberWithMutex struct {
	subscriber subscriberModel.Subscriber
	mu         sync.Mutex
}

func saveToDatabase(subscriber *SubscriberWithMutex, wg *sync.WaitGroup) *models.AppError {
	defer wg.Done()
	defer subscriber.mu.Unlock()

	subscriber.mu.Lock()

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

func SetSubscribers(subscribers *[]subscriberModel.Subscriber) *models.AppError {
	var wg sync.WaitGroup

	newSubscribers := make([]interface{}, len(*subscribers))
	for i, subscriber := range *subscribers {
		wg.Add(1)
		go saveToDatabase(&SubscriberWithMutex{
			subscriber: subscriber,
		}, &wg)
		newSubscribers[i] = subscriber
	}
	wg.Wait()

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

func GetAllSubscribers() (*[]subscriberModel.Subscriber, *models.AppError) {
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
