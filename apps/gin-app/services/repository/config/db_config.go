package db_config

import (
	"apps/gin-app/models"
	"context"
	"log/slog"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var database = "Subscribers"

var collections = map[string]string{
	"subscribers": "subscribers",
}

var SubscriberCollection *mongo.Collection

func InitializeCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database(database).Collection(collectionName)
}

func InitializeDb() (*mongo.Client, *models.AppError) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, &models.AppError{
			Err:     err,
			Message: "failed to connect to MongoDb",
			Code:    http.StatusServiceUnavailable,
		}
	}

	SubscriberCollection = InitializeCollection(client, collections["subscribers"])

	slog.Info("DB STARTED")

	return client, nil
}
