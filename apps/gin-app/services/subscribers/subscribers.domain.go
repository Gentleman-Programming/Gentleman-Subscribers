package services

import (
	"apps/gin-app/models"
	repositoryDomain "apps/gin-app/services/repository"
	subscriberModel "apps/gin-app/services/subscribers/models"
	"apps/gin-app/utils"

	"go.uber.org/fx"
)

type SubscriberService struct{}

func NewSubscribersService() SubscriberService {
	return SubscriberService{}
}

func (s *SubscriberService) ReadAndSetSubscribers() (*[]subscriberModel.Subscriber, *models.AppError) {
	subscribers, err := utils.CsvReader("files/subscribers.csv")
	if err != nil {
		return nil, err
	}

	err = repositoryDomain.SetSubscribers(subscribers)
	if err != nil {
		return nil, err
	}

	return subscribers, nil
}

func (s *SubscriberService) GetAllSubscribers() (*[]subscriberModel.Subscriber, *models.AppError) {
	subscribers, err := repositoryDomain.GetAllSubscribers()

	return subscribers, err
}

var Module = fx.Options(
	fx.Provide(NewSubscribersService),
)
