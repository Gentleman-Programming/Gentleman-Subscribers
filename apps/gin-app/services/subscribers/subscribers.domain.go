package services

import (
	"apps/gin-app/services/subscribers/models"

	"go.uber.org/fx"
)

type SubscriberService struct{}

func NewSubscribersService() SubscriberService {
	return SubscriberService{}
}

func (s *SubscriberService) GetAllSubscribers() ([]models.Subscriber, error) {
	subscribers := []models.Subscriber{
		{ID: 1, Name: "gentleman", Email: "gentlemanprogramming@gmail.com"},
		{ID: 2, Name: "juan", Email: "juan@gmail.com"},
	}

	return subscribers, nil
}

var Module = fx.Options(
	fx.Provide(NewSubscribersService),
)
